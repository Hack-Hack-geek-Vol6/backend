// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: bookmark.sql

package db

import (
	"context"
)

const createBookmark = `-- name: CreateBookmark :one
INSERT INTO bookmarks(
    hackathon_id,
    user_id
)VALUES(
    $1,$2
)RETURNING hackathon_id, user_id, create_at
`

type CreateBookmarkParams struct {
	HackathonID int32  `json:"hackathon_id"`
	UserID      string `json:"user_id"`
}

func (q *Queries) CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmarks, error) {
	row := q.db.QueryRowContext(ctx, createBookmark, arg.HackathonID, arg.UserID)
	var i Bookmarks
	err := row.Scan(&i.HackathonID, &i.UserID, &i.CreateAt)
	return i, err
}

const listBookmarkByUserID = `-- name: ListBookmarkByUserID :many
SELECT hackathon_id, user_id, create_at FROM bookmarks WHERE user_id = $1
`

func (q *Queries) ListBookmarkByUserID(ctx context.Context, userID string) ([]Bookmarks, error) {
	rows, err := q.db.QueryContext(ctx, listBookmarkByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Bookmarks{}
	for rows.Next() {
		var i Bookmarks
		if err := rows.Scan(&i.HackathonID, &i.UserID, &i.CreateAt); err != nil {
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

const removeBookmark = `-- name: RemoveBookmark :exec
DELETE FROM bookmarks WHERE user_id = $1 AND hackathon_id = $2
`

type RemoveBookmarkParams struct {
	UserID      string `json:"user_id"`
	HackathonID int32  `json:"hackathon_id"`
}

func (q *Queries) RemoveBookmark(ctx context.Context, arg RemoveBookmarkParams) error {
	_, err := q.db.ExecContext(ctx, removeBookmark, arg.UserID, arg.HackathonID)
	return err
}
