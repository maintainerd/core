package runner

import (
	"log"

	"github.com/maintainerd/core/db/seeder/v1"
	"gorm.io/gorm"
)

func RunMonolithSeeders(db *gorm.DB, appVersion string) {
	log.Println("🏃 Running seeders...")
	seeder.SeedService(db, appVersion)
	log.Println("✅ Seeding process completed.")
}
