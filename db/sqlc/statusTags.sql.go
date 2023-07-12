// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: statusTags.sql

package db

import (
	"context"
)

const getStatusTagByStatusID = `-- name: GetStatusTagByStatusID :one
SELECT status_id, status FROM status_tags WHERE status_id = $1
`

func (q *Queries) GetStatusTagByStatusID(ctx context.Context, statusID int32) (StatusTags, error) {
	row := q.db.QueryRowContext(ctx, getStatusTagByStatusID, statusID)
	var i StatusTags
	err := row.Scan(&i.StatusID, &i.Status)
	return i, err
}

const getStatusTagsByhackathonID = `-- name: GetStatusTagsByhackathonID :many
SELECT status_tags.status_id ,status_tags.status
FROM status_tags
LEFT OUTER JOIN hackathon_status_tags
ON status_tags.status_id = hackathon_status_tags.status_id
where hackathon_id = $1
`

func (q *Queries) GetStatusTagsByhackathonID(ctx context.Context, hackathonID int32) ([]StatusTags, error) {
	rows, err := q.db.QueryContext(ctx, getStatusTagsByhackathonID, hackathonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StatusTags{}
	for rows.Next() {
		var i StatusTags
		if err := rows.Scan(&i.StatusID, &i.Status); err != nil {
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

const listStatusTags = `-- name: ListStatusTags :many
SELECT status_id, status
FROM status_tags
`

func (q *Queries) ListStatusTags(ctx context.Context) ([]StatusTags, error) {
	rows, err := q.db.QueryContext(ctx, listStatusTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StatusTags{}
	for rows.Next() {
		var i StatusTags
		if err := rows.Scan(&i.StatusID, &i.Status); err != nil {
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