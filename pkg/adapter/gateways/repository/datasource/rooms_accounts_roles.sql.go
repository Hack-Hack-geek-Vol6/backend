// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: rooms_accounts_roles.sql

package repository

import (
	"context"
)

const createRoomsAccountsRoles = `-- name: CreateRoomsAccountsRoles :one
INSERT INTO rooms_accounts_roles (rooms_account_id, role_id)
VALUES ($1, $2)
RETURNING rooms_account_id, role_id
`

type CreateRoomsAccountsRolesParams struct {
	RoomsAccountID int32 `json:"rooms_account_id"`
	RoleID         int32 `json:"role_id"`
}

func (q *Queries) CreateRoomsAccountsRoles(ctx context.Context, arg CreateRoomsAccountsRolesParams) (RoomsAccountsRole, error) {
	row := q.db.QueryRowContext(ctx, createRoomsAccountsRoles, arg.RoomsAccountID, arg.RoleID)
	var i RoomsAccountsRole
	err := row.Scan(&i.RoomsAccountID, &i.RoleID)
	return i, err
}

const deleteRoomsAccountsRolesByID = `-- name: DeleteRoomsAccountsRolesByID :exec
DELETE FROM rooms_accounts_roles
WHERE rooms_account_id = $1
  AND role_id = $2
`

type DeleteRoomsAccountsRolesByIDParams struct {
	RoomsAccountID int32 `json:"rooms_account_id"`
	RoleID         int32 `json:"role_id"`
}

func (q *Queries) DeleteRoomsAccountsRolesByID(ctx context.Context, arg DeleteRoomsAccountsRolesByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteRoomsAccountsRolesByID, arg.RoomsAccountID, arg.RoleID)
	return err
}

const getRoomsAccountsRolesIDByIDs = `-- name: GetRoomsAccountsRolesIDByIDs :one
SELECT rooms_account_id
FROM rooms_accounts
WHERE room_id = $1
  AND account_id = $2
`

type GetRoomsAccountsRolesIDByIDsParams struct {
	RoomID    string `json:"room_id"`
	AccountID string `json:"account_id"`
}

func (q *Queries) GetRoomsAccountsRolesIDByIDs(ctx context.Context, arg GetRoomsAccountsRolesIDByIDsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRoomsAccountsRolesIDByIDs, arg.RoomID, arg.AccountID)
	var rooms_account_id int32
	err := row.Scan(&rooms_account_id)
	return rooms_account_id, err
}

const listRoomsAccountsRolesByID = `-- name: ListRoomsAccountsRolesByID :many
SELECT roles.role_id,
  roles.role
FROM roles
  LEFT OUTER JOIN rooms_accounts_roles ON rooms_accounts_roles.role_id = roles.role_id
WHERE rooms_accounts_roles.rooms_account_id = (
    SELECT rooms_account_id
    FROM rooms_accounts
    WHERE room_id = $1
      AND account_id = $2
  )
`

type ListRoomsAccountsRolesByIDParams struct {
	RoomID    string `json:"room_id"`
	AccountID string `json:"account_id"`
}

func (q *Queries) ListRoomsAccountsRolesByID(ctx context.Context, arg ListRoomsAccountsRolesByIDParams) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, listRoomsAccountsRolesByID, arg.RoomID, arg.AccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Role{}
	for rows.Next() {
		var i Role
		if err := rows.Scan(&i.RoleID, &i.Role); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
