package repository

import (
	"github.com/google/uuid"
	"github.com/maintainerd/core/internal/model"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Create(organization *model.Organization) error
	FindAll() ([]model.Organization, error)
	FindByUUID(organizationUUID uuid.UUID) (*model.Organization, error)
	UpdateByUUID(organizationUUID uuid.UUID, updatedOrg *model.Organization) error
	DeleteByUUID(organizationUUID uuid.UUID) error
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db}
}

func (r *organizationRepository) Create(organization *model.Organization) error {
	return r.db.Create(organization).Error
}

func (r *organizationRepository) FindAll() ([]model.Organization, error) {
	var organizations []model.Organization
	err := r.db.Find(&organizations).Error
	return organizations, err
}

func (r *organizationRepository) FindByUUID(organizationUUID uuid.UUID) (*model.Organization, error) {
	var organization model.Organization
	err := r.db.Where("organization_uuid = ?", organizationUUID).First(&organization).Error
	return &organization, err
}

func (r *organizationRepository) UpdateByUUID(organizationUUID uuid.UUID, updatedOrg *model.Organization) error {
	return r.db.Model(&model.Organization{}).
		Where("organization_uuid = ?", organizationUUID).
		Updates(updatedOrg).Error
}

func (r *organizationRepository) DeleteByUUID(organizationUUID uuid.UUID) error {
	return r.db.Where("organization_uuid = ?", organizationUUID).
		Delete(&model.Organization{}).Error
}
