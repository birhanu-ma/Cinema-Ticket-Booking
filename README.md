# Cinema Schedule & Tickets

A full-stack cinema ticket booking platform. Visitors can browse movies and showtimes; registered users can buy tickets, bookmark showtimes, and rate movies; admins manage the entire catalog — movies, directors, stars, cinemas, halls, seats, and schedules — from a dedicated admin panel.

## Tech Stack

**Frontend**
- Nuxt 4 / Vue 3
- Vite
- Vue Apollo (`@apollo/client`) for GraphQL
- vee-validate + Zod/Yup for form validation
- Tailwind CSS

**Backend**
- Go (1.22+) with Gin — powers Hasura Actions and the Chapa payment webhook
- Hasura GraphQL Engine (local, via Docker) — GraphQL API, permissions, JWT auth boundary
- PostgreSQL

**Payments**
- Chapa API (Ethiopian payment gateway)

---

## Features

### Public (no login required)
- Browse all movies and cinema schedules
- Browse by genre (`/genres` → `/genres/:id`)
- Browse by director (`/directors` → `/directors/:id`)
- Search movies by title (debounced, case-insensitive)
- Filter by genre, rating, and schedule date range
- View movie details: description, duration, director, cast, ratings, showtimes, seat availability

### Registered Users (JWT-authenticated)
- Sign up / log in
- Buy tickets for a schedule — select a hall, select one or more seats, pay via Chapa
- Bookmark a specific movie schedule (not just the movie — the showtime)
- Rate a movie (1–5 stars + optional review), one rating per user per movie, locked after submission
- View confirmed tickets ("My Tickets")
- View saved bookmarks ("My Bookmarks")
- Any of the above (buy / bookmark / rate) redirects a logged-out visitor to `/auth/signup`

### Admin Panel (separate layout, `admin` role only)
- CRUD: Directors, Stars, Cinemas, Cinema Halls, Seats
- Create Movies — multiple images with one marked as featured, one director, multiple genres, multiple stars, title/description/duration
- Create Schedules — assign a movie to a hall with price, start time, language, format; auto-generates one `schedule_seats` row per seat in that hall
- View all tickets/bookings with payment status
- Image uploads (director photos, star photos, movie images) go through a Hasura Action → Go backend → saved to local disk, served as static files
- No public admin signup — admin accounts are created directly in the database

---

## Architecture Notes

### Authentication
- JWT-based. Go backend issues tokens on `/login` and `/signup` (exposed to the frontend as Hasura Actions), embedding Hasura-required claims:
  ```json
  {
    "https://hasura.io/jwt/claims": {
      "x-hasura-default-role": "user",
      "x-hasura-allowed-roles": ["user"],
      "x-hasura-user-id": "<uuid>"
    }
  }
  ```
- Admin accounts carry `x-hasura-default-role: "admin"` instead, seeded directly via SQL/Hasura console — never through public signup.
- Frontend decodes the JWT client-side after login to route the user to `/admin/movies` or `/user/movies` accordingly.

### Hasura Actions (Go-backed)
| Action | Purpose |
|---|---|
| `login` / `signup` | Issue JWT, create user record |
| `createBooking` | Locks selected seats, creates `tickets` rows (one per seat, shared `booking_reference`) |
| `initiatePayment` | Sums the booking's total, creates a `payments` row, calls Chapa, returns a checkout URL |
| `uploadImage` | Accepts base64-encoded image, decodes and saves to local disk, returns a URL (admin-only) |

### Payment Flow
1. User selects seats → `createBooking` action locks seats and creates pending `tickets`
2. `initiatePayment` action creates a pending `payments` row and redirects to Chapa checkout
3. Chapa calls back to a Go webhook (`/payment/webhook`) — this is the **only** source of truth for payment success, not the client redirect
4. Webhook verifies the transaction with Chapa, then updates `payments.status` and all matching `tickets.status` (by shared `booking_reference`) atomically in one transaction
5. On failure, seats are released back to `is_available = true` in the same transaction

One payment can cover multiple tickets (multi-seat order) via the shared `booking_reference` — `payments.ticket_id` is kept only as an optional convenience pointer, not the real link.

### Postgres Functions, Triggers, Events, Computed Fields
The technical requirements call for demonstrating these Postgres/Hasura capabilities. Suggested approach for each, to implement as the project matures:

- **Postgres function**: a `SUM`-based function computing a schedule's remaining available seat count, exposed via Hasura as a **computed field** on `schedules` (e.g. `available_seats_count`) — avoids the client having to count `schedule_seats` manually on every render.
- **Postgres trigger**: on `payments` update (status → `success`), auto-update `tickets.status` and release/lock `schedule_seats.is_available` at the database layer as a safety net, in addition to the application-layer transaction already in the Go webhook — belt-and-suspenders against any future direct DB writes bypassing the Go service.
- **Hasura event trigger**: on `tickets` insert with `status = confirmed`, fire a webhook (could log to an audit table, or eventually send a confirmation email/SMS).
- **Generated/computed column**: a computed field on `movies` combining `ratings_aggregate.avg.rating` and count into a single "popularity score" for potential future sorting.

*(These are documented as the intended design; confirm against the actual current migration files before considering this section complete — some may already be implemented via SQL migrations not fully covered in this document.)*

### Permissions Pattern
Every user-owned table (`tickets`, `payments`, `bookmarks`, `ratings`) follows the same Hasura permission shape for the `user` role:
- **Select/Update/Delete**: row filter `{ "user_id": { "_eq": "X-Hasura-User-Id" } }`
- **Insert**: `user_id` is never client-supplied — set via a column preset reading the `X-Hasura-User-Id` session variable, preventing a user from writing data as someone else

`ratings` select is deliberately public (no row filter) so the aggregate average is visible to anonymous browsers, while insert remains user-scoped.

---

## Project Structure

```
/
├── backend/                 # Go service
│   ├── main.go
│   ├── config/
│   ├── database/
│   ├── routers/
│   └── services/
│       ├── auth.go          # signup/login
│       ├── payment.go       # createBooking, initiatePayment, webhook
│       └── upload.go        # uploadImage
├── frontend/                 # Nuxt 4 app
│   ├── pages/
│   │   ├── auth/
│   │   ├── user/             # movies, tickets, bookmarks, genres, directors
│   │   ├── admin/            # movies, directors, stars, cinemas, halls, seats, schedules, tickets
│   │   └── payment/success.vue
│   ├── components/
│   │   ├── Booking/
│   │   ├── Admin/
│   │   ├── ui/
│   │   └── Movie/
│   └── plugins/apollo.ts
└── hasura/
    ├── migrations/
    └── metadata/
```

---

## Local Setup

### Prerequisites
- Docker (for Hasura + Postgres)
- Go 1.22+
- Node.js + npm
- Hasura CLI (`hasura version` to check)
- ngrok (for exposing the local payment webhook to Chapa)

### 1. Start Hasura + Postgres
```bash
docker compose up -d
```

### 2. Apply migrations & metadata
```bash
cd hasura
hasura migrate apply --database-name default --envfile .env
hasura metadata apply --envfile .env
```

### 3. Run the Go backend
```bash
cd backend
go run main.go
```
Requires environment variables for `CHAPA_SECRET_KEY` and database connection config (see `config/`).

### 4. Expose the webhook (for local Chapa callbacks)
```bash
ngrok http 8081
```
Update the `CallbackURL` used in `initiatePayment` (ideally via an env var, not hardcoded) to the ngrok URL each time it changes.

### 5. Run the frontend
```bash
cd frontend
npm install
npm run dev
```

### 6. Create an admin user
Since there's no public admin signup, insert one directly:
```sql
UPDATE users SET role = 'admin' WHERE email = 'your-admin-email@example.com';
```
(Assuming a `role` column drives the JWT claim — adjust to match how your Go `login` handler actually derives `x-hasura-default-role`.)

---

## Known Limitations / Follow-ups
- Movies with no schedule yet are intentionally hidden from the public browse list (by design — nothing to book otherwise)
- Seats are locked at booking-creation time (not payment-confirmation time), with a `FOR UPDATE` row lock to prevent race conditions between concurrent buyers; there's currently no automatic expiry/release job for abandoned pending bookings — worth adding a scheduled cleanup
- Rating average is filtered client-side by minimum-rating filters (Postgres/Hasura can't filter on an aggregate directly without a computed view) — a Postgres function/materialized view would be a cleaner long-term fix

---

## Future Work

### Responsive UI
- The app is currently built and tested primarily for desktop viewports. Mobile/tablet breakpoints (Tailwind `sm:`/`md:`/`lg:` variants) need a full pass across:
  - Movie browse grid and filter bar (currently a horizontal scroll of filter pills — needs a mobile-friendly collapsed/drawer pattern)
  - Seat selection carousel (2×6 grid assumes desktop pointer precision; needs larger touch targets on small screens)
  - Admin panel layout (currently a fixed two-column layout with a sticky sidebar — not usable on mobile at all yet; admin is reasonably assumed desktop-only for now, but worth a decision)
  - Navigation/sidebar (needs a hamburger/drawer pattern for narrow viewports)

### Testing
- No automated tests currently exist. Plan to add:
  - **Unit tests** (Go): `services/payment.go`, `services/upload.go` — especially seat-locking logic and webhook status transitions
  - **Unit tests** (Vue): form validation logic in `vee-validate` schemas, computed properties (e.g. `currentSchedule`, `filteredMovies`)
  - **Integration tests**: full booking flow (create booking → initiate payment → webhook → ticket confirmed), using a test Chapa sandbox key
  - **E2E tests** (e.g. Playwright or Cypress): signup → browse → book → pay → view ticket, and the equivalent admin CRUD flows

### Deployment
- Currently local-only (Docker Compose for Hasura/Postgres, `go run` for the backend, `npm run dev` for the frontend). Production deployment plan:
  - **Hasura + Postgres**: move to Hasura Cloud or a managed Postgres (e.g. Neon, Supabase, RDS) instead of local Docker
  - **Go backend**: containerize with a `Dockerfile`, deploy to a host supporting long-running processes (Fly.io, Render, Railway, or a VPS) — needs a real domain for the Chapa webhook callback instead of ngrok
  - **Frontend**: `nuxt build` + deploy to Vercel/Netlify, or a Node server host if SSR is kept
  - **Environment/secrets**: move `CHAPA_SECRET_KEY`, DB credentials, and JWT secret out of local `.env` files into the hosting platform's secret manager
  - **File uploads**: local disk storage won't survive redeploys on most platforms — migrate `uploadImage` to a persistent object store (S3-compatible, e.g. Cloudflare R2 or AWS S3) before going live
  - **CORS**: set `HASURA_GRAPHQL_CORS_DOMAIN` to the real production frontend domain, not left unset/wildcard