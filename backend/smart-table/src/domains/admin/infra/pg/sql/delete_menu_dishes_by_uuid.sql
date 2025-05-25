-- name: DeleteMenuDishesByUUID :exec
--
-- args:
-- $1 - UUID[] (menu_dish_uuids)

DELETE FROM smart_table_admin.menu_dishes
WHERE uuid = ANY($1::UUID[]);
