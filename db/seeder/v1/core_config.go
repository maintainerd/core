package seeder

import (
	"log"

	"github.com/google/uuid"
	"github.com/maintainerd/core/internal/model"
	"gorm.io/gorm"
)

func SeedCoreConfigs(db *gorm.DB, appVersion string) {
	if appVersion == "" {
		log.Printf("⚠️ Skipping CoreConfig seeding: version is empty")
		return
	}

	var existing model.CoreConfig
	err := db.Where("version = ?", appVersion).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		coreConfig := model.CoreConfig{
			CoreConfigUUID: uuid.New(),
			Version:        appVersion,
			IsActive:       true,
			IsApplied:      true,
		}

		if err := db.Create(&coreConfig).Error; err != nil {
			log.Printf("❌ Failed to seed CoreConfig version '%s': %v", appVersion, err)
			return
		}

		log.Printf("✅ CoreConfig version '%s' seeded successfully", appVersion)
		return
	}

	if err != nil {
		log.Printf("❌ Error checking existing CoreConfig version '%s': %v", appVersion, err)
		return
	}

	log.Printf("⚠️ CoreConfig version '%s' already exists, skipping seeding", appVersion)
}
