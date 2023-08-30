package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type RoomUsecase interface {
	ListRooms(ctx context.Context, query domain.ListRequest) ([]domain.ListRoomResponse, error)
	GetRoom(ctx context.Context, id string) (result domain.GetRoomResponse, err error)
	CreateRoom(ctx context.Context, body domain.CreateRoomParam) (result domain.GetRoomResponse, err error)
	UpdateRoom(ctx context.Context, body domain.UpdateRoomParam) (result domain.GetRoomResponse, err error)
	DeleteRoom(ctx context.Context, query domain.DeleteRoomParam) error
	AddAccountInRoom(ctx context.Context, query domain.AddAccountInRoomParam) error
	AddChat(ctx context.Context, body domain.AddChatParams) error
	DeleteRoomAccount(ctx context.Context, body domain.DeleteRoomAccount) (err error)
}
