package model

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	OrganizationID   int        `gorm:"column:organization_id;primaryKey"`
	OrganizationUUID uuid.UUID  `gorm:"column:organization_uuid;type:uuid;not null;unique"`
	Name             string     `gorm:"column:name;type:varchar(255);not null"`
	Description      *string    `gorm:"column:description;type:text"`
	IsActive         bool       `gorm:"column:is_active;type:boolean;default:true"`
	CreatedAt        time.Time  `gorm:"column:created_at;type:timestamptz;autoCreateTime"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:timestamptz;autoUpdateTime"`
}

func (Organization) TableName() string {
	return "organizations"
}
