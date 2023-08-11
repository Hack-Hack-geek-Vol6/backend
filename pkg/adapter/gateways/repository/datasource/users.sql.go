// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package repository

import (
	"context"
	"database/sql"
	"time"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO users (
    user_id,
    email,
    hashed_password
)VALUES(
    $1,$2,$3
)RETURNING user_id, email, hashed_password, create_at, update_at, is_delete
`

type CreateUsersParams struct {
	UserID         string         `json:"user_id"`
	Email          sql.NullString `json:"email"`
	HashedPassword sql.NullString `json:"hashed_password"`
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers, arg.UserID, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.HashedPassword,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
	)
	return i, err
}

const deleteUsersByID = `-- name: DeleteUsersByID :exec
UPDATE users SET is_delete = $1 WHERE user_id = $2
`

type DeleteUsersByIDParams struct {
	IsDelete bool   `json:"is_delete"`
	UserID   string `json:"user_id"`
}

func (q *Queries) DeleteUsersByID(ctx context.Context, arg DeleteUsersByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteUsersByID, arg.IsDelete, arg.UserID)
	return err
}

const getUsersByEmail = `-- name: GetUsersByEmail :one
SELECT user_id, email, hashed_password, create_at, update_at, is_delete FROM users WHERE email = $1
`

func (q *Queries) GetUsersByEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsersByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.HashedPassword,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
	)
	return i, err
}

const getUsersByID = `-- name: GetUsersByID :one
SELECT user_id, email, hashed_password, create_at, update_at, is_delete FROM users WHERE user_id = $1
`

func (q *Queries) GetUsersByID(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsersByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.HashedPassword,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
	)
	return i, err
}

const updateUsersByID = `-- name: UpdateUsersByID :one
UPDATE users
SET email = $1,
    hashed_password = $2,
    update_at = $3
WHERE user_id = $4
RETURNING user_id, email, hashed_password, create_at, update_at, is_delete
`

type UpdateUsersByIDParams struct {
	Email          sql.NullString `json:"email"`
	HashedPassword sql.NullString `json:"hashed_password"`
	UpdateAt       time.Time      `json:"update_at"`
	UserID         string         `json:"user_id"`
}

func (q *Queries) UpdateUsersByID(ctx context.Context, arg UpdateUsersByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUsersByID,
		arg.Email,
		arg.HashedPassword,
		arg.UpdateAt,
		arg.UserID,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.HashedPassword,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
	)
	return i, err
}