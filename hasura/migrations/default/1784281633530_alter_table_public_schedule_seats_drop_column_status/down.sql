alter table "public"."schedule_seats" alter column "status" drop not null;
alter table "public"."schedule_seats" add column "status" text;
