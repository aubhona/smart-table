-- name: FetchRestaurantsByUUID :many
--
-- args:
-- $1 - UUID[] (restaurant_uuids)

SELECT jsonb_build_object(
    'restaurant', to_jsonb(r),
    'dishes', (
        COALESCE(
            (
                 SELECT jsonb_agg(to_jsonb(d) ORDER BY d.name)
                 FROM smart_table_admin.dishes d
                 WHERE d.restaurant_uuid = r.uuid
             ),
            '[]'::jsonb
        )
    ),
    'owner', (
        SELECT
            to_jsonb(u)
        FROM
            smart_table_admin.users u
        WHERE
            u.uuid = r.owner_uuid
    )
) AS restaurant_data
FROM smart_table_admin.restaurants r
WHERE r.uuid = ANY($1::UUID[])
ORDER BY r.uuid;
