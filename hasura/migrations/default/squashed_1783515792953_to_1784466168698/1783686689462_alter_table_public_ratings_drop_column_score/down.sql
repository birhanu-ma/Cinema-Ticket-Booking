alter table "public"."ratings" alter column "score" drop not null;
alter table "public"."ratings" add column "score" int4;
