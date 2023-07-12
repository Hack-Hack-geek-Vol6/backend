// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: hackathon.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createHackathon = `-- name: CreateHackathon :one
INSERT INTO hackathons (
    name,
    icon,
    description,
    link,
    expired,
    start_date,
    term
  )
VALUES(
    $1,$2,$3,$4,$5,$6,$7
  )
RETURNING hackathon_id, name, icon, description, link, expired, start_date, term
`

type CreateHackathonParams struct {
	Name        string         `json:"name"`
	Icon        sql.NullString `json:"icon"`
	Description string         `json:"description"`
	Link        string         `json:"link"`
	Expired     time.Time      `json:"expired"`
	StartDate   time.Time      `json:"start_date"`
	Term        int32          `json:"term"`
}

func (q *Queries) CreateHackathon(ctx context.Context, arg CreateHackathonParams) (Hackathons, error) {
	row := q.db.QueryRowContext(ctx, createHackathon,
		arg.Name,
		arg.Icon,
		arg.Description,
		arg.Link,
		arg.Expired,
		arg.StartDate,
		arg.Term,
	)
	var i Hackathons
	err := row.Scan(
		&i.HackathonID,
		&i.Name,
		&i.Icon,
		&i.Description,
		&i.Link,
		&i.Expired,
		&i.StartDate,
		&i.Term,
	)
	return i, err
}

const getHackathonByID = `-- name: GetHackathonByID :one
SELECT hackathon_id, name, icon, description, link, expired, start_date, term
FROM hackathons
WHERE hackathon_id = $1
`

func (q *Queries) GetHackathonByID(ctx context.Context, hackathonID int32) (Hackathons, error) {
	row := q.db.QueryRowContext(ctx, getHackathonByID, hackathonID)
	var i Hackathons
	err := row.Scan(
		&i.HackathonID,
		&i.Name,
		&i.Icon,
		&i.Description,
		&i.Link,
		&i.Expired,
		&i.StartDate,
		&i.Term,
	)
	return i, err
}

const listHackathons = `-- name: ListHackathons :many
SELECT hackathon_id, name, icon, description, link, expired, start_date, term
FROM hackathons
WHERE expired > $1
ORDER BY hackathon_id
LIMIT $2 OFFSET $3
`

type ListHackathonsParams struct {
	Expired time.Time `json:"expired"`
	Limit   int32     `json:"limit"`
	Offset  int32     `json:"offset"`
}

func (q *Queries) ListHackathons(ctx context.Context, arg ListHackathonsParams) ([]Hackathons, error) {
	rows, err := q.db.QueryContext(ctx, listHackathons, arg.Expired, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Hackathons{}
	for rows.Next() {
		var i Hackathons
		if err := rows.Scan(
			&i.HackathonID,
			&i.Name,
			&i.Icon,
			&i.Description,
			&i.Link,
			&i.Expired,
			&i.StartDate,
			&i.Term,
		); err != nil {
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
