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
	ServiceHandler      *handler.ServiceHandler
}

func NewApp(db *gorm.DB) *App {
	// repository
	organizationRepo := repository.NewOrganizationRepository(db)
	serviceRepo := repository.NewServiceRepository(db)

	// service
	organizationService := service.NewOrganizationService(organizationRepo)
	serviceService := service.NewServiceService(serviceRepo)

	// handler
	organizationHandler := handler.NewOrganizationHandler(organizationService)
	serviceHandler := handler.NewServiceHandler(serviceService)

	return &App{
		DB:                  db,
		OrganizationHandler: organizationHandler,
		ServiceHandler:      serviceHandler,
	}
}
