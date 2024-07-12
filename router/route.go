package router

import (
	"RjsConfigService/ent"
	"RjsConfigService/handler"
	"RjsConfigService/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, dbClient *ent.Client) {
	api := app.Group("/api", logger.New())

	setUpConfigRoutes(api, dbClient)
	api.Get("/health", handler.HealthEndpoint())
}

func setUpConfigRoutes(api fiber.Router, dbClient *ent.Client) {
	config := api.Group("/v1/config")
	config.Use(middleware.IsExpired())

	config.Get("/:AppName", handler.FetchConfigEndpoint(dbClient))
	config.Post("/:AppName", handler.CreateConfigEndpoint(dbClient))
	config.Post("/:AppName/activate/:Version", handler.ActivateConfigEndpoint(dbClient))
}
