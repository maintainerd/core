package app

import (
	"github.com/maintainerd/core/internal/handler"
	"github.com/maintainerd/core/internal/repository"
	"github.com/maintainerd/core/internal/service"
	"gorm.io/gorm"
)

type App struct {
	DB                  *gorm.DB
	OrganizationHandler *handler.OrganizationHandler
}

func NewApp(db *gorm.DB) *App {
	// repository
	organizationRepo := repository.NewOrganizationRepository(db)

	// service
	organizationService := service.NewOrganizationService(organizationRepo)

	// handler
	organizationHandler := handler.NewOrganizationHandler(organizationService)

	return &App{
		DB:                  db,
		OrganizationHandler: organizationHandler,
	}
}
