package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/maintainerd/core/internal/model"
)

type OrganizationDTO struct {
	OrganizationUUID uuid.UUID  `json:"organization_uuid"`
	Name             string     `json:"name"`
	Description      *string    `json:"description,omitempty"`
	IsActive         bool       `json:"is_active"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

func ToOrganizationDTO(organization *model.Organization) OrganizationDTO {
	return OrganizationDTO{
		OrganizationUUID: organization.OrganizationUUID,
		Name:             organization.Name,
		Description:      organization.Description,
		IsActive:         organization.IsActive,
		CreatedAt:        organization.CreatedAt,
		UpdatedAt:        organization.UpdatedAt,
	}
}
