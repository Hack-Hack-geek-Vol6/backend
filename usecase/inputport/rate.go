package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/domain"
	repository "github.com/hackhack-Geek-vol6/backend/gateways/repository/datasource"
)

type RateUsecase interface {
	CreateRateEntry(ctx context.Context, body repository.CreateRateParams) (repository.RateEntry, error)
	ListRateEntry(ctx context.Context, id string, query domain.ListRateParams) ([]repository.RateEntry, error)
}
