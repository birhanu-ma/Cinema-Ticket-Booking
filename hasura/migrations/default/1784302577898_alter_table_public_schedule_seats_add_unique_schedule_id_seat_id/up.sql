alter table "public"."schedule_seats" drop constraint "schedule_seats_schedule_id_key";
alter table "public"."schedule_seats" add constraint "schedule_seats_schedule_id_seat_id_key" unique ("schedule_id", "seat_id");
