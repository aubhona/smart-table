-- name: UpsertEmployees :exec
--
-- args:
-- $1 - JSONB ([]PgEmployee)

INSERT INTO smart_table_admin.employees (
    user_uuid,
    place_uuid,
    role,
    active,
    created_at,
    updated_at
)
SELECT
    input.user_uuid,
    input.place_uuid,
    input.role,
    input.active,
    input.created_at,
    input.updated_at
FROM jsonb_to_recordset($1::jsonb) AS input(
    user_uuid       UUID,
    place_uuid      UUID,
    role            TEXT,
    active          BOOLEAN,
    created_at      TIMESTAMPTZ,
    updated_at      TIMESTAMPTZ
)
ON CONFLICT ON CONSTRAINT uniq_employee_place  DO NOTHING;