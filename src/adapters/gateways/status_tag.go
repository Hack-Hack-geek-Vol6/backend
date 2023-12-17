package gateways

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/datastructure/models"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type StatusTagGateway struct {
	db *gorm.DB
}

func NewStatusTagGateway(db *gorm.DB) dai.StatusTagDai {
	return &StatusTagGateway{
		db: db,
	}
}

func (stg *StatusTagGateway) Create(ctx context.Context, statusTag *models.StatusTag) (id int64, err error) {
	result := stg.db.Create(statusTag)
	if result.Error != nil {
		return 0, result.Error
	}

	return statusTag.ID, nil
}

func (stg *StatusTagGateway) FindAll(ctx context.Context) (statusTags []*models.StatusTag, err error) {
	result := stg.db.Find(&statusTags)
	if result.Error != nil {
		return nil, result.Error
	}

	return statusTags, nil
}

func (stg *StatusTagGateway) FindById(ctx context.Context, id int64) (statusTag *models.StatusTag, err error) {
	result := stg.db.First(&statusTag, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return statusTag, nil
}

func (stg *StatusTagGateway) Update(ctx context.Context, statusTag *models.StatusTag) (id int64, err error) {
	result := stg.db.Save(statusTag)
	if result.Error != nil {
		return 0, result.Error
	}

	return statusTag.ID, nil
}
