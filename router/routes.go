package router

import (
	"config-service/ent"
	"config-service/handler"
	"config-service/middleware"
	"time"

	"github.com/apialerts/apialerts-go"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/patrickmn/go-cache"
)

func SetupRoutes(server *fuego.Server, dbClient *ent.Client, apiAlertsClient apialerts.Client, jwtSecretKey string) {
	api := fuego.Group(server, "/api")

	setUpConfigRoutes(api, dbClient, apiAlertsClient, jwtSecretKey)
	setUpUserRoutes(api, dbClient, apiAlertsClient, jwtSecretKey)
}

func setUpConfigRoutes(api *fuego.Server, dbClient *ent.Client, apiAlertsClient apialerts.Client, jwtSecretKey string) {
	configCache := cache.New(5*time.Minute, 10*time.Minute)
	handlers := &handler.ConfigResources{DbClient: dbClient, ApiAlertsClient: apiAlertsClient, Cache: configCache}
	config := fuego.Group(api, "/v1/config",
		fuego.OptionSecurity(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)
	fuego.Use(config, middleware.IsExpired(jwtSecretKey))

	fuego.Get(config, "/apps", handlers.GetConfigApps)
	fuego.Get(config, "/{AppName}/all", handlers.GetAllConfigsForApp)
	fuego.Get(config, "/{AppName}", handlers.GetConfig)
	fuego.Post(config, "/{AppName}", handlers.CreateConfig)
	fuego.Post(config, "/{AppName}/activate/{Version}", handlers.ActivateConfig)
}

func setUpUserRoutes(api *fuego.Server, dbClient *ent.Client, apiAlertsClient apialerts.Client, jwtSecretKey string) {
	handlers := &handler.UserResources{DbClient: dbClient, ApiAlertsClient: apiAlertsClient}

	user := fuego.Group(api, "/v1/user",
		fuego.OptionSecurity(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)
	fuego.Use(user, middleware.IsExpired(jwtSecretKey))

	fuego.Post(user, "/{AppName}", handlers.CreateUser)
	fuego.Get(user, "/", handlers.GetAllUsers)
}
