CREATE EXTENSION IF NOT EXISTS pgcrypto;

SET check_function_bodies = false;

CREATE TABLE public.bookmarks (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    schedule_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.cinema_halls (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    cinema_id uuid NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    capacity integer NOT NULL
);
CREATE TABLE public.cinemas (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    location text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    update_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.directors (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    bio text,
    photo_url text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    update_at timestamp with time zone DEFAULT now()
);
CREATE TABLE public.genres (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.movie_genres (
    movie_id uuid NOT NULL,
    genre_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.movie_images (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    movie_id uuid NOT NULL,
    image_url text NOT NULL,
    is_featured boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.movie_stars (
    movie_id uuid NOT NULL,
    star_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.movies (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    director_id uuid,
    duration integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.payments (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    ticket_id uuid,
    transaction_ref text,
    status text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    amount integer,
    paid_at time with time zone,
    currency text,
    payment_method text,
    booking_reference text
);
CREATE TABLE public.ratings (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    movie_id uuid NOT NULL,
    rating integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    review text,
    CONSTRAINT ratings_rating_check CHECK (((rating >= 1) AND (rating <= 5)))
);
CREATE TABLE public.schedule_seats (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    schedule_id uuid DEFAULT gen_random_uuid() NOT NULL,
    seat_id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    is_available boolean,
    locked_by uuid,
    locked_at timestamp with time zone
);
CREATE TABLE public.schedules (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    movies_id uuid,
    hall_id uuid,
    price numeric(10,2) NOT NULL,
    start_time timestamp with time zone NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    language text,
    format text,
    end_time timestamp with time zone
);
CREATE TABLE public.seats (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    hall_id uuid DEFAULT gen_random_uuid() NOT NULL,
    seat_number integer NOT NULL,
    row_name text NOT NULL,
    type text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.stars (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    bio text,
    photo_url text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    update_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE TABLE public.tickets (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    schedule_id uuid,
    status text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    booking_reference text,
    schedule_seat_id uuid DEFAULT gen_random_uuid()
);
CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    role character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY public.bookmarks ADD CONSTRAINT bookmarks_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.bookmarks ADD CONSTRAINT bookmarks_user_id_schedule_id_key UNIQUE (user_id, schedule_id);
ALTER TABLE ONLY public.cinema_halls ADD CONSTRAINT cinema_halls_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.cinemas ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.directors ADD CONSTRAINT directors_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.genres ADD CONSTRAINT genres_name_key UNIQUE (name);
ALTER TABLE ONLY public.genres ADD CONSTRAINT genres_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.movie_genres ADD CONSTRAINT movie_genres_pkey PRIMARY KEY (movie_id, genre_id);
ALTER TABLE ONLY public.movie_images ADD CONSTRAINT movie_images_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.movie_stars ADD CONSTRAINT movie_stars_pkey PRIMARY KEY (movie_id, star_id);
ALTER TABLE ONLY public.movies ADD CONSTRAINT movies_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.payments ADD CONSTRAINT payments_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.ratings ADD CONSTRAINT ratings_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.ratings ADD CONSTRAINT ratings_user_id_movie_id_key UNIQUE (user_id, movie_id);
ALTER TABLE ONLY public.schedule_seats ADD CONSTRAINT schedule_seats_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.schedule_seats ADD CONSTRAINT schedule_seats_schedule_id_seat_id_key UNIQUE (schedule_id, seat_id);
ALTER TABLE ONLY public.schedules ADD CONSTRAINT schedules_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.seats ADD CONSTRAINT seats_hall_id_seat_number_row_name_key UNIQUE (hall_id, seat_number, row_name);
ALTER TABLE ONLY public.seats ADD CONSTRAINT seats_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.stars ADD CONSTRAINT stars_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.tickets ADD CONSTRAINT tickets_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.users ADD CONSTRAINT users_email_key UNIQUE (email);
ALTER TABLE ONLY public.users ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.bookmarks ADD CONSTRAINT bookmarks_schedule_id_fkey FOREIGN KEY (schedule_id) REFERENCES public.schedules(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.bookmarks ADD CONSTRAINT bookmarks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.cinema_halls ADD CONSTRAINT cinema_halls_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.movie_genres ADD CONSTRAINT movie_genres_genre_id_fkey FOREIGN KEY (genre_id) REFERENCES public.genres(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.movie_genres ADD CONSTRAINT movie_genres_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.movie_images ADD CONSTRAINT movie_images_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.movie_stars ADD CONSTRAINT movie_stars_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.movie_stars ADD CONSTRAINT movie_stars_star_id_fkey FOREIGN KEY (star_id) REFERENCES public.stars(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.movies ADD CONSTRAINT movies_director_id_fkey FOREIGN KEY (director_id) REFERENCES public.directors(id) ON DELETE SET NULL;
ALTER TABLE ONLY public.payments ADD CONSTRAINT payments_ticket_id_fkey FOREIGN KEY (ticket_id) REFERENCES public.tickets(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.payments ADD CONSTRAINT payments_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.ratings ADD CONSTRAINT ratings_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.ratings ADD CONSTRAINT ratings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.schedule_seats ADD CONSTRAINT schedule_seats_schedule_id_fkey FOREIGN KEY (schedule_id) REFERENCES public.schedules(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.schedule_seats ADD CONSTRAINT schedule_seats_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seats(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.schedules ADD CONSTRAINT schedules_hall_id_fkey FOREIGN KEY (hall_id) REFERENCES public.cinema_halls(id) ON UPDATE CASCADE ON DELETE RESTRICT;
ALTER TABLE ONLY public.schedules ADD CONSTRAINT schedules_movies_id_fkey FOREIGN KEY (movies_id) REFERENCES public.movies(id) ON DELETE SET NULL;
ALTER TABLE ONLY public.seats ADD CONSTRAINT seats_hall_id_fkey FOREIGN KEY (hall_id) REFERENCES public.cinema_halls(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.tickets ADD CONSTRAINT tickets_schedule_id_fkey FOREIGN KEY (schedule_id) REFERENCES public.schedules(id) ON DELETE SET NULL;
ALTER TABLE ONLY public.tickets ADD CONSTRAINT tickets_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL;

-- fix_schema_gaps: the three fixes from earlier in this cleanup pass
ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_schedule_seat_id_fkey
    FOREIGN KEY (schedule_seat_id)
    REFERENCES public.schedule_seats(id)
    ON DELETE SET NULL;

ALTER TABLE public.payments
    ALTER COLUMN paid_at TYPE timestamp with time zone
    USING (CURRENT_DATE + paid_at);

ALTER TABLE public.schedule_seats ALTER COLUMN schedule_id DROP DEFAULT;
ALTER TABLE public.schedule_seats ALTER COLUMN seat_id DROP DEFAULT;
ALTER TABLE public.seats ALTER COLUMN hall_id DROP DEFAULT;
ALTER TABLE public.tickets ALTER COLUMN schedule_seat_id DROP DEFAULT;
