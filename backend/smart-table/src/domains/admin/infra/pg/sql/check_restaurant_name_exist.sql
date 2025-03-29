-- name: CheckNameExist :one
--
-- args:
-- $1 - TEXT (name)

SELECT EXISTS (
    SELECT 1
    FROM 
        smart_table_admin.restaurants r
    WHERE 
        r.name = $1::TEXT
) AS restaurant_exists;
