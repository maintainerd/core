package service

import (
	"github.com/google/uuid"
	"github.com/maintainerd/core/internal/model"
	"github.com/maintainerd/core/internal/repository"
)

type OrganizationService interface {
	Create(organization *model.Organization) error
	GetAll() ([]model.Organization, error)
	GetByUUID(organizationUUID uuid.UUID) (*model.Organization, error)
	UpdateByUUID(organizationUUID uuid.UUID, updatedOrganization *model.Organization) error
	DeleteByUUID(organizationUUID uuid.UUID) error
}

type organizationService struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return &organizationService{repo}
}

func (s *organizationService) Create(organization *model.Organization) error {
	organization.OrganizationUUID = uuid.New()
	return s.repo.Create(organization)
}

func (s *organizationService) GetAll() ([]model.Organization, error) {
	return s.repo.FindAll()
}

func (s *organizationService) GetByUUID(organizationUUID uuid.UUID) (*model.Organization, error) {
	return s.repo.FindByUUID(organizationUUID)
}

func (s *organizationService) UpdateByUUID(organizationUUID uuid.UUID, updatedOrganization *model.Organization) error {
	return s.repo.UpdateByUUID(organizationUUID, updatedOrganization)
}

func (s *organizationService) DeleteByUUID(organizationUUID uuid.UUID) error {
	return s.repo.DeleteByUUID(organizationUUID)
}
