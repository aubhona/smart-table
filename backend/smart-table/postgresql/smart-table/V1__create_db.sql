CREATE TABLE IF NOT EXISTS "smart-table.customers" (
   "uuid" UUID PRIMARY KEY NOT NULL,
   "tg_id" TEXT NOT NULL,
   "tg_login" TEXT NOT NULL UNIQUE,
   "avatar_link" TEXT NOT NULL,
   "chat_id" TEXT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "smart-table.orders" (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "room_code" TEXT NOT NULL,
    "table_id" TEXT NOT NULL,
    "customers_uuid" UUID[] NOT NULL,
    "host_user_uuid" UUID NOT NULL,
    "status" TEXT NOT NULL,
    "resolution" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "smart-table.items" (
   "uuid" UUID PRIMARY KEY NOT NULL,
   "order_uuid" UUID NOT NULL,
   "comment" TEXT,
   "status" TEXT NOT NULL,
   "resolution" TEXT,
   "name" TEXT NOT NULL,
   "description" TEXT NOT NULL,
   "picture_link" TEXT NOT NULL,
   "weight" INT NOT NULL,
   "category" TEXT NOT NULL,
   "price" DECIMAL NOT NULL,
   "customer_uuid" UUID NOT NULL,
   "dish_uuid" UUID NOT NULL,
   "is_draft" BOOLEAN NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "smart-table.dishes" (
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

CREATE TABLE IF NOT EXISTS "smart-table.menu-dishes" (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "dish_uuid" UUID NOT NULL,
    "place_uuid" UUID NOT NULL,
    "price" DECIMAL NOT NULL,
    "exist" BOOLEAN DEFAULT TRUE NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "smart-table.places" (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "rest_uuid" UUID NOT NULL,
    "address" TEXT NOT NULL,
    "opening_time" TIME NOT NULL,
    "closing_time" TIME NOT NULL,
    "table_count" INT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "smart-table.restaurants" (
     "uuid" UUID PRIMARY KEY NOT NULL,
     "name" TEXT NOT NULL,
     "owner_uuid" UUID NOT NULL,
     "created_at" TIMESTAMPTZ NOT NULL,
     "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "smart-table.users" (
   "uuid" UUID PRIMARY KEY NOT NULL,
   "avatar_link" TEXT NOT NULL,
   "login" TEXT NOT NULL,
   "tg_login" TEXT NOT NULL,
   "name" TEXT NOT NULL,
   "phone" TEXT NOT NULL,
   "password_hash" TEXT NOT NULL,
   "chat_id" TEXT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "smart-table.staff" (
   "user_uuid" UUID PRIMARY KEY NOT NULL,
   "place_uuid" UUID NOT NULL,
   "role" TEXT NOT NULL,
   "active" BOOLEAN DEFAULT TRUE NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
DECLARE tbl TEXT;
BEGIN
FOR tbl IN SELECT tablename FROM pg_tables WHERE schemaname = 'public' LOOP
        EXECUTE format(
            'CREATE TRIGGER trigger_%s_updated_at
            BEFORE UPDATE ON %s
            FOR EACH ROW
            EXECUTE FUNCTION set_timestamp();', tbl, tbl
        );
END LOOP;
END $$;

ALTER TABLE "smart-table.orders" ADD CONSTRAINT fk_orders_customers FOREIGN KEY ("customers_uuid") REFERENCES "smart-table.customers" ("uuid");
ALTER TABLE "smart-table.items" ADD CONSTRAINT fk_items_orders FOREIGN KEY ("order_uuid") REFERENCES "smart-table.orders" ("uuid");
ALTER TABLE "smart-table.items" ADD CONSTRAINT fk_items_customers FOREIGN KEY ("customer_uuid") REFERENCES "smart-table.customers" ("uuid");
ALTER TABLE "smart-table.items" ADD CONSTRAINT fk_items_dishes FOREIGN KEY ("dish_uuid") REFERENCES "smart-table.dishes" ("uuid");
ALTER TABLE "smart-table.dishes" ADD CONSTRAINT fk_dishes_restaurants FOREIGN KEY ("rest_uuid") REFERENCES "smart-table.restaurants" ("uuid");
ALTER TABLE "smart-table.menu" ADD CONSTRAINT fk_menu_dishes FOREIGN KEY ("dish_uuid") REFERENCES "smart-table.dishes" ("uuid");
ALTER TABLE "smart-table.menu" ADD CONSTRAINT fk_menu_places FOREIGN KEY ("place_uuid") REFERENCES "smart-table.places" ("uuid");
ALTER TABLE "smart-table.places" ADD CONSTRAINT fk_places_restaurants FOREIGN KEY ("rest_uuid") REFERENCES "smart-table.restaurants" ("uuid");
ALTER TABLE "smart-table.restaurants" ADD CONSTRAINT fk_restaurants_users FOREIGN KEY ("owner_uuid") REFERENCES "smart-table.users" ("uuid");
ALTER TABLE "smart-table.staff" ADD CONSTRAINT fk_staff_users FOREIGN KEY ("user_uuid") REFERENCES "smart-table.users" ("uuid");
ALTER TABLE "smart-table.staff" ADD CONSTRAINT fk_staff_places FOREIGN KEY ("place_uuid") REFERENCES "smart-table.places" ("uuid");
