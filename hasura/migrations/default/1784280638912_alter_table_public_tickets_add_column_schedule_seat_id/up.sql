CREATE EXTENSION IF NOT EXISTS pgcrypto;
alter table "public"."tickets" add column "schedule_seat_id" uuid
 null default gen_random_uuid();
