package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type FollowUsecase interface {
	CreateFollow(ctx context.Context, body repository.CreateFollowsParams) (result domain.FollowResponse, err error)
	RemoveFollow(ctx context.Context, body repository.DeleteFollowsParams) error
	GetFollowByToID(ctx context.Context, ID string) (result []domain.FollowResponse, err error)
}