package router

import (
	"RjsConfigService/handler"
	"RjsConfigService/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	setUpConfigFetchRoutes(api)
	setUpAdminRoutes(api)
}

func setUpConfigFetchRoutes(api fiber.Router) {
	configFetch := api.Group("/v1/config")
	configFetch.Use(middleware.IsExpired())

	configFetch.Get("/:AppName", handler.FetchConfigEndpoint)
}

func setUpAdminRoutes(api fiber.Router) {
	admin := api.Group("/v1/admin")
	admin.Use(middleware.IsExpired())
}
