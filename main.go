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
	// env := config.Config("ENV")

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
	go database.CreateUserAccounts(client, jwtSecret)

	apiAlertsClient := alert.Connect()

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	router.SetupRoutes(s, client, *apiAlertsClient, jwtSecret)

	s.Run()
}
