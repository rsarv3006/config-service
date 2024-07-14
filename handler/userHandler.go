package handler

import (
	"RjsConfigService/auth"
	"RjsConfigService/config"
	"RjsConfigService/ent"
	"context"
	"log"

	"github.com/apialerts/apialerts-go"
	"github.com/gofiber/fiber/v2"
)

func CreateAppUserEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("CreateAppUserEndpoint")
		currentUser := c.Locals("currentUser").(*ent.User)

		if currentUser.Role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
				"error":   nil,
			})
		}

		appName := c.Params("AppName")

		if appName == "" {
			return sendBadRequestResponse(c, nil, "Invalid request app name must be defined")
		}

		log.Println("appName: " + appName)
		user, err := dbClient.User.
			Create().
			SetAppName(appName).
			SetRole("user").
			Save(context.Background())

		if err != nil {
			log.Println(err)
			return sendBadRequestResponse(c, err, "Error creating user")
		}

		log.Println("user: ")

		jwtSecret := config.Config("JWT_SECRET")
		log.Println("jwtSecret: " + jwtSecret)
		token, err := auth.GenerateJWTFromSecret(user, jwtSecret)

		if err != nil {
			log.Println(err)
			return sendBadRequestResponse(c, err, "Error creating user")
		}

		log.Println("token: " + token)
		apiAlertsClient := c.Locals("ApiAlertsClient").(*apialerts.Client)
		log.Println("apiAlertsClient: ")
		go apiAlertsClient.Send("User created for app: "+appName,
			[]string{appName},
			"v1/config/"+appName)

		log.Println("apiAlertsClient.Send")
		return c.JSON(fiber.Map{
			"message": "Success",
			"data":    user,
			"token":   token,
		})
	}
}
