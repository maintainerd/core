package main

import (
	"log"
	"os"

	"github.com/maintainerd/core/config"
	"github.com/maintainerd/core/db/runner"
	"github.com/maintainerd/core/internal/app"
	"github.com/maintainerd/core/internal/repository"
	"github.com/maintainerd/core/internal/route"
	"github.com/maintainerd/core/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	connString := config.GetDBConnectionString()

	appVersion := os.Getenv("APP_VERSION")
	appMode := os.Getenv("APP_MODE")

	if appMode == "mono" {
		serviceRepository := repository.NewServiceRepository(db)
		serviceService := service.NewServiceService(serviceRepository)

		_, err := serviceService.GetByName("core")
		if err != nil {
			// Run default seeders
			if appVersion == "v0.0.1" {
				runner.RunMonolithMigrations(connString)
				runner.RunMonolithSeeders(db, appVersion)
			}
		}
	}

	application := app.NewApp(db)
	r := gin.Default()

	route.RegisterRoute(r, &route.HandlerCollection{
		OrganizationHandler: application.OrganizationHandler,
		ServiceHandler:      application.ServiceHandler,
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
