package model

import (
	"time"

	"github.com/google/uuid"
)

type CoreConfig struct {
	CoreConfigID   int        `gorm:"column:core_config_id;primaryKey"`
	CoreConfigUUID uuid.UUID  `gorm:"column:core_config_uuid;type:uuid;not null;unique"`
	Version        string     `gorm:"column:version;type:varchar(20);not null"`
	IsActive       bool       `gorm:"column:is_active;type:boolean;default:true"`
	IsApplied      bool       `gorm:"column:is_applied;type:boolean;default:true"`
	CreatedAt      time.Time  `gorm:"column:created_at;type:timestamptz;autoCreateTime"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;type:timestamptz;autoUpdateTime"`
}

func (CoreConfig) TableName() string {
	return "core_config"
}
