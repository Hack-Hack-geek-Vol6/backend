// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateAccountFramework(ctx context.Context, arg CreateAccountFrameworkParams) (AccountFrameworks, error)
	CreateAccountTags(ctx context.Context, arg CreateAccountTagsParams) (AccountTags, error)
	CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmarks, error)
	CreateFollow(ctx context.Context, arg CreateFollowParams) (Follows, error)
	CreateHackathon(ctx context.Context, arg CreateHackathonParams) (Hackathons, error)
	CreateHackathonStatusTag(ctx context.Context, arg CreateHackathonStatusTagParams) (HackathonStatusTags, error)
	CreatePastWorkFrameworks(ctx context.Context, arg CreatePastWorkFrameworksParams) (PastWorkFrameworks, error)
	CreatePastWorkTag(ctx context.Context, arg CreatePastWorkTagParams) (PastWorkTags, error)
	CreatePastWorks(ctx context.Context, arg CreatePastWorksParams) (PastWorks, error)
	CreateRate(ctx context.Context, arg CreateRateParams) (RateEntries, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Rooms, error)
	CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccounts, error)
	DeleteAccountFrameworksByUserID(ctx context.Context, userID string) error
	DeleteAccounttagsByUserID(ctx context.Context, userID string) error
	DeleteFrameworksByID(ctx context.Context, frameworkID int32) error
	DeleteHackathonByID(ctx context.Context, hackathonID int32) error
	DeleteHackathonStatusTagsByHackathonID(ctx context.Context, hackathonID int32) error
	DeletePastWorkFrameworksByOpus(ctx context.Context, opus int32) error
	DeletePastWorkTagsByOpus(ctx context.Context, opus int32) error
	DeleteStatusTagByStatusID(ctx context.Context, statusID int32) error
	DeleteTechTagByID(ctx context.Context, techTagID int32) error
	GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error)
	GetAccountByID(ctx context.Context, userID string) (GetAccountByIDRow, error)
	GetFrameworksByID(ctx context.Context, frameworkID int32) (Frameworks, error)
	GetHackathonByID(ctx context.Context, hackathonID int32) (Hackathons, error)
	GetHackathonStatusTagsByHackathonID(ctx context.Context, hackathonID int32) ([]HackathonStatusTags, error)
	GetLocateByID(ctx context.Context, locateID int32) (Locates, error)
	GetPastWorkFrameworksByOpus(ctx context.Context, opus int32) ([]PastWorkFrameworks, error)
	GetPastWorksByOpus(ctx context.Context, opus int32) (PastWorks, error)
	GetRoomsAccountsByRoomID(ctx context.Context, roomID uuid.UUID) ([]GetRoomsAccountsByRoomIDRow, error)
	GetRoomsByID(ctx context.Context, roomID uuid.UUID) (Rooms, error)
	GetStatusTagByStatusID(ctx context.Context, statusID int32) (StatusTags, error)
	GetStatusTagsByHackathonID(ctx context.Context, hackathonID int32) ([]StatusTags, error)
	GetTechTagByID(ctx context.Context, techTagID int32) (TechTags, error)
	ListAccountFrameworksByUserID(ctx context.Context, userID string) ([]ListAccountFrameworksByUserIDRow, error)
	ListAccountTagsByUserID(ctx context.Context, userID string) ([]ListAccountTagsByUserIDRow, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]ListAccountsRow, error)
	ListBookmarkByUserID(ctx context.Context, userID string) ([]Bookmarks, error)
	ListFollowByToUserID(ctx context.Context, toUserID string) ([]Follows, error)
	ListFrameworks(ctx context.Context, limit int32) ([]Frameworks, error)
	ListHackathons(ctx context.Context, arg ListHackathonsParams) ([]Hackathons, error)
	ListLocates(ctx context.Context) ([]Locates, error)
	ListPastWorkTagsByOpus(ctx context.Context, opus int32) ([]PastWorkTags, error)
	ListPastWorks(ctx context.Context, arg ListPastWorksParams) ([]ListPastWorksRow, error)
	ListRate(ctx context.Context, arg ListRateParams) ([]RateEntries, error)
	ListRoom(ctx context.Context, limit int32) ([]Rooms, error)
	ListStatusTags(ctx context.Context) ([]StatusTags, error)
	ListTechTag(ctx context.Context) ([]TechTags, error)
	RemoveAccountInRoom(ctx context.Context, arg RemoveAccountInRoomParams) error
	RemoveFollow(ctx context.Context, arg RemoveFollowParams) error
	SoftDeleteAccount(ctx context.Context, userID string) (Accounts, error)
	SoftDeleteRoomByID(ctx context.Context, roomID uuid.UUID) (Rooms, error)
	SoftRemoveBookmark(ctx context.Context, arg SoftRemoveBookmarkParams) (Bookmarks, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accounts, error)
	UpdateFrameworksByID(ctx context.Context, arg UpdateFrameworksByIDParams) (Frameworks, error)
	UpdateRateByUserID(ctx context.Context, arg UpdateRateByUserIDParams) (Accounts, error)
	UpdateRoomByID(ctx context.Context, arg UpdateRoomByIDParams) (Rooms, error)
	UpdateStatusTagByStatusID(ctx context.Context, status string) (StatusTags, error)
	UpdateTechTagByID(ctx context.Context, language string) (TechTags, error)
}

var _ Querier = (*Queries)(nil)
