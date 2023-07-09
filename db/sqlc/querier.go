// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreataAccountTags(ctx context.Context, arg CreataAccountTagsParams) (AccountTags, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateHackathon(ctx context.Context, arg CreateHackathonParams) (Hackathons, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Rooms, error)
	CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccounts, error)
	CreateRoomsTechTag(ctx context.Context, arg CreateRoomsTechTagParams) (RoomsTechTags, error)
	GetAccount(ctx context.Context, userID string) (GetAccountRow, error)
	GetAccountAuth(ctx context.Context, userID string) (GetAccountAuthRow, error)
	GetAccountTags(ctx context.Context, userID string) ([]GetAccountTagsRow, error)
	GetHackathon(ctx context.Context, hackathonID int32) (Hackathons, error)
	GetLocate(ctx context.Context, locateID int32) (Locates, error)
	GetRoom(ctx context.Context, roomID uuid.UUID) (Rooms, error)
	GetRoomsAccounts(ctx context.Context, roomID uuid.UUID) ([]GetRoomsAccountsRow, error)
	GetRoomsTechTags(ctx context.Context, roomID uuid.UUID) ([]GetRoomsTechTagsRow, error)
	GetTechTag(ctx context.Context, techTagID int32) (TechTags, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]ListAccountsRow, error)
	ListHackathons(ctx context.Context, arg ListHackathonsParams) ([]Hackathons, error)
	ListLocates(ctx context.Context) ([]Locates, error)
	ListRoom(ctx context.Context, limit int32) ([]Rooms, error)
	ListTechTag(ctx context.Context) ([]TechTags, error)
}

var _ Querier = (*Queries)(nil)
