-- name: CreateRoomsAccountsRoles :one
INSERT INTO rooms_accounts_roles (rooms_account_id, role_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListRoomsAccountsRolesByRoomID :many
SELECT *
FROM rooms_accounts_roles
WHERE rooms_account_id = $1
LIMIT $2 OFFSET $3;

-- name: DeleteRoomsAccountsRolesByRoomID :exec
DELETE FROM rooms_accounts_roles
WHERE rooms_account_id = $1;