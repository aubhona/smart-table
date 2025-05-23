-- +goose Up
-- +goose StatementBegin
BEGIN;
CREATE INDEX IF NOT EXISTS idx_places_restaurant_uuid
  ON smart_table_admin.places (restaurant_uuid);

CREATE UNIQUE INDEX IF NOT EXISTS idx_places_address_restaurant_uuid
  ON smart_table_admin.places (address, restaurant_uuid);

CREATE INDEX IF NOT EXISTS idx_restaurants_owner_uuid
  ON smart_table_admin.restaurants (owner_uuid);

CREATE UNIQUE INDEX IF NOT EXISTS idx_restaurants_name
  ON smart_table_admin.restaurants (name);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_login
  ON smart_table_admin.users (login);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_tg_login
  ON smart_table_admin.users (tg_login);

CREATE INDEX IF NOT EXISTS idx_employees_user_uuid
  ON smart_table_admin.employees (user_uuid);

CREATE INDEX IF NOT EXISTS idx_employees_place_uuid
  ON smart_table_admin.employees (place_uuid);

CREATE INDEX IF NOT EXISTS idx_dishes_restaurant_uuid
  ON smart_table_admin.dishes (restaurant_uuid);

CREATE INDEX IF NOT EXISTS idx_menu_dishes_place_uuid
  ON smart_table_admin.menu_dishes (place_uuid);

CREATE INDEX IF NOT EXISTS idx_menu_dishes_dish_uuid
  ON smart_table_admin.menu_dishes (dish_uuid); 

END;
-- +goose StatementEnd
