package service

import (
	"github.com/google/uuid"
	"github.com/maintainerd/core/internal/model"
	"github.com/maintainerd/core/internal/repository"
)

type CoreConfigService interface {
	Create(role *model.CoreConfig) error
	GetLatestConfig() (*model.CoreConfig, error)
}

type coreConfigService struct {
	repo repository.CoreConfigRepository
}

func (s *coreConfigService) Create(coreConfig *model.CoreConfig) error {
	coreConfig.CoreConfigUUID = uuid.New()
	return s.repo.Create(coreConfig)
}

func NewCoreConfigService(repo repository.CoreConfigRepository) CoreConfigService {
	return &coreConfigService{repo}
}

func (s *coreConfigService) GetLatestConfig() (*model.CoreConfig, error) {
	return s.repo.FetchLatest()
}
