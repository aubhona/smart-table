-- name: FetchPlacesByUUIDForUpdate :many
--
-- args:
-- $1 - UUID[] (place_uuids)

SELECT jsonb_build_object(
    'place', to_jsonb(p),
    'restaurant', (
        SELECT jsonb_build_object(
           'restaurant', to_jsonb(r),
           'dishes', (
               SELECT COALESCE(
                  jsonb_agg(to_jsonb(d)),
                  '[]'::jsonb
               )
               FROM smart_table_admin.dishes d
               WHERE d.restaurant_uuid = r.uuid
           ),
           'owner', (
               SELECT to_jsonb(u)
               FROM smart_table_admin.users u
               WHERE u.uuid = r.owner_uuid
           )
        )
        FROM smart_table_admin.restaurants r
        WHERE r.uuid = p.restaurant_uuid
    ),
    'employees', (
        SELECT COALESCE(
            jsonb_agg(
                jsonb_build_object(
                    'employee', to_jsonb(e),
                    'user', (
                        SELECT to_jsonb(u)
                        FROM smart_table_admin.users u
                        WHERE u.uuid = e.user_uuid
                    )
                )
            ),
            '[]'::jsonb
        )
        FROM smart_table_admin.employees e
        WHERE e.place_uuid = p.uuid
    ),
    'menu_dishes', (
        SELECT COALESCE(
           jsonb_agg(
               jsonb_build_object(
                   'menu_dish', to_jsonb(md),
                   'dish', (
                       SELECT to_jsonb(d)
                       FROM smart_table_admin.dishes d
                       WHERE d.uuid = md.dish_uuid
                   )
               )
           ),
           '[]'::jsonb
        )
        FROM smart_table_admin.menu_dishes md
        WHERE md.place_uuid = p.uuid
    )
) AS place_data
FROM smart_table_admin.places p
WHERE p.uuid = ANY($1::UUID[])
FOR UPDATE;
