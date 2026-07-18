alter table "public"."seats" drop constraint "seats_hall_id_seat_number_key";
alter table "public"."seats" add constraint "seats_hall_id_seat_number_row_name_key" unique ("hall_id", "seat_number", "row_name");
