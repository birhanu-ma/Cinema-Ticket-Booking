alter table "public"."payments" alter column "currency" drop not null;
alter table "public"."payments" add column "currency" int4;
