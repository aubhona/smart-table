-- name: CheckLoginOrTgLoginExist :one
--
-- args:
-- $1 - TEXT (login)
-- $2 - TEXT (tg_login)

SELECT EXISTS (
    SELECT 1
    FROM 
        smart_table_admin.users u 
    WHERE 
        u.login = $1::TEXT
        OR u.tg_login = $2::TEXT
) AS user_exists;
