alter table "public"."bookmarks" add constraint "bookmarks_user_id_schedule_id_key" unique ("user_id", "schedule_id");
