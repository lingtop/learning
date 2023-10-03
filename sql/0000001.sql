-- DROP TABLE IF EXISTS "public"."sign_set_applications";

CREATE SEQUENCE IF NOT EXISTS sign_set_applications_id_seq;

CREATE TABLE "public"."sign_set_applications" (
    "id" int8 NOT NULL DEFAULT nextval('sign_set_applications_id_seq'::regclass),
    "name" text NOT NULL,
    "secret" text NOT NULL,
    "callback_url" text NOT NULL,
    "updated_time" int8 NOT NULL,
    "created_time" int8 NOT NULL,
    PRIMARY KEY ("id")
);

