alter table "public"."ratings" add constraint "ratings_user_id_movie_id_key" unique ("user_id", "movie_id");
