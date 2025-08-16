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

func SetupRoutes(server *fuego.Server, dbClient *ent.Client, apiAlertsClient apialerts.Client, jwtSecretKey string, v2SecretKey string) {
	api := fuego.Group(server, "/api")

	setUpConfigRoutes(api, dbClient, apiAlertsClient, jwtSecretKey, v2SecretKey)
	setUpUserRoutes(api, dbClient, apiAlertsClient, jwtSecretKey, v2SecretKey)
}

func setUpConfigRoutes(api *fuego.Server, dbClient *ent.Client, apiAlertsClient apialerts.Client, jwtSecretKey string, v2SecretKey string) {
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

	// V2
	configV2 := fuego.Group(api, "/v2/config",
		fuego.OptionSecurity(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)
	fuego.Use(configV2, middleware.IsExpired(v2SecretKey))

	fuego.Get(configV2, "/apps", handlers.GetConfigApps)
	fuego.Get(configV2, "/{AppName}/all", handlers.GetAllConfigsForApp)
	fuego.Get(configV2, "/{AppName}", handlers.GetConfig)
	fuego.Post(configV2, "/{AppName}", handlers.CreateConfig)
	fuego.Post(configV2, "/{AppName}/activate/{Version}", handlers.ActivateConfig)
}

func setUpUserRoutes(api *fuego.Server, dbClient *ent.Client, apiAlertsClient apialerts.Client, jwtSecretKey string, v2SecretKey string) {
	handlers := &handler.UserResources{DbClient: dbClient, ApiAlertsClient: apiAlertsClient}

	user := fuego.Group(api, "/v1/user",
		fuego.OptionSecurity(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)
	fuego.Use(user, middleware.IsExpired(jwtSecretKey))
	fuego.Get(user, "/", handlers.GetAllUsers)

	// V2
	userV2 := fuego.Group(api, "/v2/user",
		fuego.OptionSecurity(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)
	fuego.Use(userV2, middleware.IsExpired(v2SecretKey))

	fuego.Post(userV2, "/{AppName}", handlers.CreateUser)
	fuego.Get(userV2, "/", handlers.GetAllUsers)

}
