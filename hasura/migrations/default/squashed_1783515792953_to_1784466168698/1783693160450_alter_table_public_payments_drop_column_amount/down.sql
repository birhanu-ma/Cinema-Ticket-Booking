alter table "public"."payments" alter column "amount" drop not null;
alter table "public"."payments" add column "amount" int4;
