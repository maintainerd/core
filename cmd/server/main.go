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
		coreConfigRepository := repository.NewCoreConfigRepository(db)
		coreConfigService := service.NewCoreConfigService(coreConfigRepository)

		_, err := coreConfigService.GetLatestConfig()
		if err != nil {
			// Run default seeders
			if appVersion == "v0.0.1" {
				runner.RunDefaultMigrations(connString)
				runner.RunDefaultSeeders(db, appVersion)
			}
		}
	}

	application := app.NewApp(db)
	r := gin.Default()

	route.Registerroute(r, &route.HandlerCollection{
		OrganizationHandler: application.OrganizationHandler,
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
