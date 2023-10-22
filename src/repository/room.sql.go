// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: room.sql

package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const closeRoomByID = `-- name: CloseRoomByID :one
UPDATE rooms
SET is_closing = true
WHERE room_id = $1
RETURNING room_id, hackathon_id, title, description, member_limit, include_rate, create_at, update_at, is_delete, is_closing
`

func (q *Queries) CloseRoomByID(ctx context.Context, roomID string) (Room, error) {
	row := q.db.QueryRow(ctx, closeRoomByID, roomID)
	var i Room
	err := row.Scan(
		&i.RoomID,
		&i.HackathonID,
		&i.Title,
		&i.Description,
		&i.MemberLimit,
		&i.IncludeRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.IsClosing,
	)
	return i, err
}

const createRooms = `-- name: CreateRooms :one
INSERT INTO rooms (
        room_id,
        hackathon_id,
        title,
        description,
        member_limit,
        include_rate,
        is_closing
    )
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING room_id, hackathon_id, title, description, member_limit, include_rate, create_at, update_at, is_delete, is_closing
`

type CreateRoomsParams struct {
	RoomID      string      `json:"room_id"`
	HackathonID int32       `json:"hackathon_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	MemberLimit int32       `json:"member_limit"`
	IncludeRate bool        `json:"include_rate"`
	IsClosing   pgtype.Bool `json:"is_closing"`
}

func (q *Queries) CreateRooms(ctx context.Context, arg CreateRoomsParams) (Room, error) {
	row := q.db.QueryRow(ctx, createRooms,
		arg.RoomID,
		arg.HackathonID,
		arg.Title,
		arg.Description,
		arg.MemberLimit,
		arg.IncludeRate,
		arg.IsClosing,
	)
	var i Room
	err := row.Scan(
		&i.RoomID,
		&i.HackathonID,
		&i.Title,
		&i.Description,
		&i.MemberLimit,
		&i.IncludeRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.IsClosing,
	)
	return i, err
}

const deleteRoomsByID = `-- name: DeleteRoomsByID :one
UPDATE rooms
SET is_delete = true
WHERE room_id = $1
RETURNING room_id, hackathon_id, title, description, member_limit, include_rate, create_at, update_at, is_delete, is_closing
`

func (q *Queries) DeleteRoomsByID(ctx context.Context, roomID string) (Room, error) {
	row := q.db.QueryRow(ctx, deleteRoomsByID, roomID)
	var i Room
	err := row.Scan(
		&i.RoomID,
		&i.HackathonID,
		&i.Title,
		&i.Description,
		&i.MemberLimit,
		&i.IncludeRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.IsClosing,
	)
	return i, err
}

const getRoomsByID = `-- name: GetRoomsByID :one
SELECT room_id, hackathon_id, title, description, member_limit, include_rate, create_at, update_at, is_delete, is_closing
FROM rooms
WHERE room_id = $1 AND is_delete = false
`

func (q *Queries) GetRoomsByID(ctx context.Context, roomID string) (Room, error) {
	row := q.db.QueryRow(ctx, getRoomsByID, roomID)
	var i Room
	err := row.Scan(
		&i.RoomID,
		&i.HackathonID,
		&i.Title,
		&i.Description,
		&i.MemberLimit,
		&i.IncludeRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.IsClosing,
	)
	return i, err
}

const listRooms = `-- name: ListRooms :many
SELECT room_id, hackathon_id, title, description, member_limit, include_rate, create_at, update_at, is_delete, is_closing
FROM rooms
WHERE is_delete = false
ORDER BY create_at DESC
LIMIT $1 OFFSET $2
`

type ListRoomsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRooms(ctx context.Context, arg ListRoomsParams) ([]Room, error) {
	rows, err := q.db.Query(ctx, listRooms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Room{}
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.RoomID,
			&i.HackathonID,
			&i.Title,
			&i.Description,
			&i.MemberLimit,
			&i.IncludeRate,
			&i.CreateAt,
			&i.UpdateAt,
			&i.IsDelete,
			&i.IsClosing,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRoomsByID = `-- name: UpdateRoomsByID :one
UPDATE rooms
SET hackathon_id = $1,
    title = $2,
    description = $3,
    member_limit = $4,
    update_at = $5,
    is_closing = $6
WHERE room_id = $7
RETURNING room_id, hackathon_id, title, description, member_limit, include_rate, create_at, update_at, is_delete, is_closing
`

type UpdateRoomsByIDParams struct {
	HackathonID int32       `json:"hackathon_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	MemberLimit int32       `json:"member_limit"`
	UpdateAt    time.Time   `json:"update_at"`
	IsClosing   pgtype.Bool `json:"is_closing"`
	RoomID      string      `json:"room_id"`
}

func (q *Queries) UpdateRoomsByID(ctx context.Context, arg UpdateRoomsByIDParams) (Room, error) {
	row := q.db.QueryRow(ctx, updateRoomsByID,
		arg.HackathonID,
		arg.Title,
		arg.Description,
		arg.MemberLimit,
		arg.UpdateAt,
		arg.IsClosing,
		arg.RoomID,
	)
	var i Room
	err := row.Scan(
		&i.RoomID,
		&i.HackathonID,
		&i.Title,
		&i.Description,
		&i.MemberLimit,
		&i.IncludeRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.IsClosing,
	)
	return i, err
}