-- name: FetchUserByLoginOrTgLogin :one
--
-- args:
-- $1 - TEXT (login)
-- $2 - TEXT (tg_login)

SELECT
    to_jsonb(u)
FROM
    smart_table_admin.users u
WHERE
    u.login = $1::TEXT
        OR u.tg_login = $2::TEXT;
