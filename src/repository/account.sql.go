// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: account.sql

package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkAccount = `-- name: CheckAccount :one
SELECT
    count(*)
FROM
    accounts
WHERE 
    account_id = $1 AND email = $2
`

type CheckAccountParams struct {
	AccountID string `json:"account_id"`
	Email     string `json:"email"`
}

func (q *Queries) CheckAccount(ctx context.Context, arg CheckAccountParams) (int64, error) {
	row := q.db.QueryRow(ctx, checkAccount, arg.AccountID, arg.Email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAccounts = `-- name: CreateAccounts :one
INSERT INTO
    accounts (
        account_id,
        email,
        username,
        icon,
        explanatory_text,
        locate_id,
        rate,
        character,
        show_locate,
        show_rate
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10
    ) RETURNING account_id, username, icon, explanatory_text, locate_id, rate, character, show_locate, show_rate, create_at, update_at, is_delete, email, twitter_link, github_link, discord_link
`

type CreateAccountsParams struct {
	AccountID       string      `json:"account_id"`
	Email           string      `json:"email"`
	Username        string      `json:"username"`
	Icon            pgtype.Text `json:"icon"`
	ExplanatoryText pgtype.Text `json:"explanatory_text"`
	LocateID        int32       `json:"locate_id"`
	Rate            int32       `json:"rate"`
	Character       pgtype.Int4 `json:"character"`
	ShowLocate      bool        `json:"show_locate"`
	ShowRate        bool        `json:"show_rate"`
}

func (q *Queries) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccounts,
		arg.AccountID,
		arg.Email,
		arg.Username,
		arg.Icon,
		arg.ExplanatoryText,
		arg.LocateID,
		arg.Rate,
		arg.Character,
		arg.ShowLocate,
		arg.ShowRate,
	)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.Character,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.Email,
		&i.TwitterLink,
		&i.GithubLink,
		&i.DiscordLink,
	)
	return i, err
}

const deleteAccounts = `-- name: DeleteAccounts :one
UPDATE
    accounts
SET
    is_delete = true
WHERE
    account_id = $1 RETURNING account_id, username, icon, explanatory_text, locate_id, rate, character, show_locate, show_rate, create_at, update_at, is_delete, email, twitter_link, github_link, discord_link
`

func (q *Queries) DeleteAccounts(ctx context.Context, accountID string) (Account, error) {
	row := q.db.QueryRow(ctx, deleteAccounts, accountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.Character,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.Email,
		&i.TwitterLink,
		&i.GithubLink,
		&i.DiscordLink,
	)
	return i, err
}

const getAccountsByEmail = `-- name: GetAccountsByEmail :one
SELECT
    account_id, username, icon, explanatory_text, locate_id, rate, character, show_locate, show_rate, create_at, update_at, is_delete, email, twitter_link, github_link, discord_link
FROM
    accounts
WHERE
    email = $1 AND is_delete = false
`

func (q *Queries) GetAccountsByEmail(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountsByEmail, email)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.Character,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.Email,
		&i.TwitterLink,
		&i.GithubLink,
		&i.DiscordLink,
	)
	return i, err
}

const getAccountsByID = `-- name: GetAccountsByID :one
SELECT
    account_id, username, icon, explanatory_text, locate_id, rate, character, show_locate, show_rate, create_at, update_at, is_delete, email, twitter_link, github_link, discord_link
FROM
    accounts
WHERE
    account_id = $1 AND is_delete = false
`

func (q *Queries) GetAccountsByID(ctx context.Context, accountID string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountsByID, accountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.Character,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.Email,
		&i.TwitterLink,
		&i.GithubLink,
		&i.DiscordLink,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT
    account_id, username, icon, explanatory_text, locate_id, rate, character, show_locate, show_rate, create_at, update_at, is_delete, email, twitter_link, github_link, discord_link
FROM
    accounts
WHERE
    is_delete = false
ORDER BY
    rate DESC
LIMIT
    $1 OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.AccountID,
			&i.Username,
			&i.Icon,
			&i.ExplanatoryText,
			&i.LocateID,
			&i.Rate,
			&i.Character,
			&i.ShowLocate,
			&i.ShowRate,
			&i.CreateAt,
			&i.UpdateAt,
			&i.IsDelete,
			&i.Email,
			&i.TwitterLink,
			&i.GithubLink,
			&i.DiscordLink,
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

const updateAccounts = `-- name: UpdateAccounts :one
UPDATE
    accounts
SET
    username = $2,
    icon = $3,
    explanatory_text = $4,
    locate_id = $5,
    rate = $6,
    character = $7,
    show_locate = $8,
    show_rate = $9,
    update_at = $10,
    twitter_link = $11,
    github_link = $12,
    discord_link = $13
WHERE
    account_id = $1 RETURNING account_id, username, icon, explanatory_text, locate_id, rate, character, show_locate, show_rate, create_at, update_at, is_delete, email, twitter_link, github_link, discord_link
`

type UpdateAccountsParams struct {
	AccountID       string      `json:"account_id"`
	Username        string      `json:"username"`
	Icon            pgtype.Text `json:"icon"`
	ExplanatoryText pgtype.Text `json:"explanatory_text"`
	LocateID        int32       `json:"locate_id"`
	Rate            int32       `json:"rate"`
	Character       pgtype.Int4 `json:"character"`
	ShowLocate      bool        `json:"show_locate"`
	ShowRate        bool        `json:"show_rate"`
	UpdateAt        time.Time   `json:"update_at"`
	TwitterLink     pgtype.Text `json:"twitter_link"`
	GithubLink      pgtype.Text `json:"github_link"`
	DiscordLink     pgtype.Text `json:"discord_link"`
}

func (q *Queries) UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccounts,
		arg.AccountID,
		arg.Username,
		arg.Icon,
		arg.ExplanatoryText,
		arg.LocateID,
		arg.Rate,
		arg.Character,
		arg.ShowLocate,
		arg.ShowRate,
		arg.UpdateAt,
		arg.TwitterLink,
		arg.GithubLink,
		arg.DiscordLink,
	)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.Character,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.Email,
		&i.TwitterLink,
		&i.GithubLink,
		&i.DiscordLink,
	)
	return i, err
}

const updateRateByID = `-- name: UpdateRateByID :one
UPDATE
    accounts
SET    
    rate = $2,
    update_at = $3
WHERE
    account_id = $1 RETURNING account_id, username, icon, explanatory_text, locate_id, rate, character, show_locate, show_rate, create_at, update_at, is_delete, email, twitter_link, github_link, discord_link
`

type UpdateRateByIDParams struct {
	AccountID string    `json:"account_id"`
	Rate      int32     `json:"rate"`
	UpdateAt  time.Time `json:"update_at"`
}

func (q *Queries) UpdateRateByID(ctx context.Context, arg UpdateRateByIDParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateRateByID, arg.AccountID, arg.Rate, arg.UpdateAt)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.Character,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
		&i.UpdateAt,
		&i.IsDelete,
		&i.Email,
		&i.TwitterLink,
		&i.GithubLink,
		&i.DiscordLink,
	)
	return i, err
}