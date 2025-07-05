package repository

import (
	"github.com/maintainerd/core/internal/model"
	"gorm.io/gorm"
)

type CoreConfigRepository interface {
	Create(coreConfig *model.CoreConfig) error
	FetchLatest() (*model.CoreConfig, error)
}

type coreConfigRepository struct {
	db *gorm.DB
}

func NewCoreConfigRepository(db *gorm.DB) CoreConfigRepository {
	return &coreConfigRepository{db}
}

func (r *coreConfigRepository) Create(coreConfig *model.CoreConfig) error {
	return r.db.Create(coreConfig).Error
}

func (r *coreConfigRepository) FetchLatest() (*model.CoreConfig, error) {
	var coreConfig model.CoreConfig
	err := r.db.Order("created_at DESC").First(&coreConfig).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &coreConfig, nil
}
