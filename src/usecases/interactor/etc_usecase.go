package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
)

type etcUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewEtcUsercase(store transaction.Store, timeout time.Duration) inputport.EtcUsecase {
	return &etcUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (eu *etcUsecase) GetFramework(ctx context.Context) ([]repository.Framework, error) {
	fmt.Println(eu.contextTimeout)
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListFrameworks(ctx)
}

func (eu *etcUsecase) GetLocat(ctx context.Context) ([]repository.Locate, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListLocates(ctx)
}

func (eu *etcUsecase) GetTechTag(ctx context.Context) ([]repository.TechTag, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListTechTags(ctx)
}

func (eu *etcUsecase) GetStatusTag(ctx context.Context) ([]repository.StatusTag, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListStatusTags(ctx)
}

func (eu *etcUsecase) ListRoles(ctx context.Context) ([]repository.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListRoles(ctx)
}