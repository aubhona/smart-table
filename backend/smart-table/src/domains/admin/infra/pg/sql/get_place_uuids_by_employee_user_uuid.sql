-- name: GetPlaceUUIDsByEmployeeUserUUID :many
--
-- args:
-- $1 - UUID (user_uuid)

SELECT
    e.place_uuid
FROM
    smart_table_admin.employees e
WHERE
    e.user_uuid = $1::UUID;
