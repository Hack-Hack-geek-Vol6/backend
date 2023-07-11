// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: rooms_accounts.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createRoomsAccounts = `-- name: CreateRoomsAccounts :one
INSERT INTO rooms_accounts (
    user_id,
    room_id,
    is_owner
)VALUES(
    $1,$2,$3
)RETURNING user_id, room_id, is_owner
`

type CreateRoomsAccountsParams struct {
	UserID  string    `json:"user_id"`
	RoomID  uuid.UUID `json:"room_id"`
	IsOwner bool      `json:"is_owner"`
}

func (q *Queries) CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccounts, error) {
	row := q.db.QueryRowContext(ctx, createRoomsAccounts, arg.UserID, arg.RoomID, arg.IsOwner)
	var i RoomsAccounts
	err := row.Scan(&i.UserID, &i.RoomID, &i.IsOwner)
	return i, err
}

const getRoomsAccounts = `-- name: GetRoomsAccounts :many
SELECT 
    accounts.user_id,
    accounts.username,  
    accounts.icon
FROM 
    rooms_accounts
LEFT OUTER JOIN 
    accounts 
ON 
    rooms_accounts.user_id = accounts.user_id 
WHERE 
    rooms_accounts.room_id = $1
`

type GetRoomsAccountsRow struct {
	UserID   sql.NullString `json:"user_id"`
	Username sql.NullString `json:"username"`
	Icon     []byte         `json:"icon"`
}

func (q *Queries) GetRoomsAccounts(ctx context.Context, roomID uuid.UUID) ([]GetRoomsAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getRoomsAccounts, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRoomsAccountsRow{}
	for rows.Next() {
		var i GetRoomsAccountsRow
		if err := rows.Scan(&i.UserID, &i.Username, &i.Icon); err != nil {
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

const removeAccountInRoom = `-- name: RemoveAccountInRoom :one
DELETE FROM rooms_accounts WHERE room_id = $1 AND user_id = $2 RETURNING user_id, room_id, is_owner
`

type RemoveAccountInRoomParams struct {
	RoomID uuid.UUID `json:"room_id"`
	UserID string    `json:"user_id"`
}

func (q *Queries) RemoveAccountInRoom(ctx context.Context, arg RemoveAccountInRoomParams) (RoomsAccounts, error) {
	row := q.db.QueryRowContext(ctx, removeAccountInRoom, arg.RoomID, arg.UserID)
	var i RoomsAccounts
	err := row.Scan(&i.UserID, &i.RoomID, &i.IsOwner)
	return i, err
}