package seeder

import (
	"log"
	"time"

	"github.com/maintainerd/auth/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []model.Role{
		{
			RoleUUID:    uuid.New(),
			Name:        "super-admin",
			Description: strPtr("Super Admin"),
			IsDefault:   false,
			CreatedAt:   time.Now(),
		},
		{
			RoleUUID:    uuid.New(),
			Name:        "admin",
			Description: strPtr("Admin"),
			IsDefault:   false,
			CreatedAt:   time.Now(),
		},
		{
			RoleUUID:    uuid.New(),
			Name:        "registered",
			Description: strPtr("Registered"),
			IsDefault:   true,
			CreatedAt:   time.Now(),
		},
	}

	for _, role := range roles {
		var existing model.Role
		if err := db.Where("name = ?", role.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					log.Printf("❌ Failed to seed role %s: %v\n", role.Name, err)
				} else {
					log.Printf("✅ Role %s seeded successfully\n", role.Name)
				}
			} else {
				log.Printf("❌ Error checking role %s: %v\n", role.Name, err)
			}
		} else {
			log.Printf("⚠️ Role %s already exists, skipping\n", role.Name)
		}
	}
}

func strPtr(s string) *string {
	return &s
}
