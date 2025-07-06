package seeder

import (
	"log"

	"github.com/google/uuid"
	"github.com/maintainerd/core/internal/model"
	"gorm.io/gorm"
)

func SeedService(db *gorm.DB, appVersion string) {
	if appVersion == "" {
		log.Printf("⚠️ Skipping Service seeding: version is empty")
		return
	}

	var existing model.Service
	err := db.Where("service_name = ?", "default").First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		service := model.Service{
			ServiceUUID: uuid.New(),
			ServiceName: "core",
			DisplayName: "Core Service",
			Description: "Core system service",
			ServiceType: "default",
			Version:     appVersion,
			IsActive:    true,
			IsDefault:   true,
		}

		if err := db.Create(&service).Error; err != nil {
			log.Printf("❌ Failed to seed Default Service version '%s': %v", appVersion, err)
			return
		}

		log.Printf("✅ Default Service version '%s' seeded successfully", appVersion)
		return
	}

	if err != nil {
		log.Printf("❌ Error checking existing Default Service: %v", err)
		return
	}

	log.Printf("⚠️ Default Service already exists, skipping seeding")
}
