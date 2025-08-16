package main

import (
	"config-service/alert"
	"config-service/config"
	"config-service/database"
	"config-service/router"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/rs/cors"
)

func main() {
	jwtSecret := config.Config("JWT_SECRET")
	jwtSecretV2 := config.Config("JWT_SECRET_V2")

	s := fuego.NewServer(
		fuego.WithAddr(":3000"),
		fuego.WithSecurity(map[string]*openapi3.SecuritySchemeRef{
			"bearerAuth": &openapi3.SecuritySchemeRef{
				Value: openapi3.NewSecurityScheme().
					WithType("http").
					WithScheme("bearer").
					WithBearerFormat("JWT").
					WithDescription("Enter your JWT token in the format: Bearer <token>"),
			},
		}),
		fuego.WithGlobalMiddlewares(cors.AllowAll().Handler),
	)

	client := database.Connect()
	go database.CreateUserAccounts(client, jwtSecretV2)

	apiAlertsClient := alert.Connect()

	router.SetupRoutes(s, client, *apiAlertsClient, jwtSecret, jwtSecretV2)

	s.Run()
}
