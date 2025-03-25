-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE SCHEMA IF NOT EXISTS smart_table_admin;

CREATE TABLE IF NOT EXISTS smart_table_admin.users (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "login" TEXT NOT NULL CHECK ("login" ~ '^[a-zA-Z][a-zA-Z0-9_]{4,31}$'),
    "tg_id" TEXT NOT NULL,
    "tg_login" TEXT NOT NULL CHECK ("tg_login" ~ '^[a-zA-Z][a-zA-Z0-9_]{4,31}$'),
    "chat_id" TEXT NOT NULL,
    "first_name" TEXT NOT NULL CHECK ("first_name" ~ '^[A-Za-z\\-\\s]+$'),
    "last_name" TEXT NOT NULL CHECK ("last_name" ~ '^[A-Za-z\\-\\s]+$'),
    "password_hash" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);


CREATE TABLE IF NOT EXISTS smart_table_admin.dishes (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "weight" INT NOT NULL,
    "picture_link" TEXT NOT NULL,
    "rest_uuid" UUID NOT NULL,
    "category" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS smart_table_admin.menu_dishes (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "dish_uuid" UUID NOT NULL,
    "place_uuid" UUID NOT NULL,
    "price" DECIMAL NOT NULL,
    "exist" BOOLEAN DEFAULT TRUE NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS smart_table_admin.places (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "rest_uuid" UUID NOT NULL,
    "address" TEXT NOT NULL,
    "opening_time" TIME NOT NULL,
    "closing_time" TIME NOT NULL,
    "table_count" INT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS smart_table_admin.restaurants (
     "uuid" UUID PRIMARY KEY NOT NULL,
     "name" TEXT NOT NULL,
     "owner_uuid" UUID NOT NULL,
     "created_at" TIMESTAMPTZ NOT NULL,
     "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS smart_table_admin.staff (
   "user_uuid" UUID PRIMARY KEY NOT NULL,
   "place_uuid" UUID NOT NULL,
   "role" TEXT NOT NULL,
   "active" BOOLEAN DEFAULT TRUE NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL
);

ALTER TABLE smart_table_admin.dishes ADD CONSTRAINT fk_dishes_restaurants FOREIGN KEY ("rest_uuid") REFERENCES smart_table_admin.restaurants ("uuid");
ALTER TABLE smart_table_admin.menu_dishes ADD CONSTRAINT fk_menu_dishes FOREIGN KEY ("dish_uuid") REFERENCES smart_table_admin.dishes ("uuid");
ALTER TABLE smart_table_admin.menu_dishes ADD CONSTRAINT fk_menu_places FOREIGN KEY ("place_uuid") REFERENCES smart_table_admin.places ("uuid");
ALTER TABLE smart_table_admin.places ADD CONSTRAINT fk_places_restaurants FOREIGN KEY ("rest_uuid") REFERENCES smart_table_admin.restaurants ("uuid");
ALTER TABLE smart_table_admin.restaurants ADD CONSTRAINT fk_restaurants_users FOREIGN KEY ("owner_uuid") REFERENCES smart_table_admin.users ("uuid");
ALTER TABLE smart_table_admin.staff ADD CONSTRAINT fk_staff_users FOREIGN KEY ("user_uuid") REFERENCES smart_table_admin.users ("uuid");
ALTER TABLE smart_table_admin.staff ADD CONSTRAINT fk_staff_places FOREIGN KEY ("place_uuid") REFERENCES smart_table_admin.places ("uuid");

COMMIT;
-- +goose StatementEnd
