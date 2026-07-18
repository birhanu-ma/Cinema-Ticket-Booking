alter table "public"."tickets" alter column "seat_number" drop not null;
alter table "public"."tickets" add column "seat_number" int4;
