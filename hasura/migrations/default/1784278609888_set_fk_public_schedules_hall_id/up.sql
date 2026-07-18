alter table "public"."schedules"
  add constraint "schedules_hall_id_fkey"
  foreign key ("hall_id")
  references "public"."cinema_halls"
  ("id") on update cascade on delete restrict;
