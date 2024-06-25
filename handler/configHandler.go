package handler

import (
	"RjsConfigService/ent"
	"RjsConfigService/ent/appconfig"
	"RjsConfigService/model"
	"context"
	"log"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/apialerts/apialerts-go"
	"github.com/gofiber/fiber/v2"
)

func FetchConfigEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		appName := c.Params("AppName")

		if appName == "" {
			return sendBadRequestResponse(c, nil, "AppName is required")
		}

		currentUser := c.Locals("currentUser").(*ent.User)

		if currentUser.Role != "admin" && currentUser.AppName != appName {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
				"error":   nil,
			})
		}

		config, err := getCurrentConfigForApp(appName, dbClient)

		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"message": "Error fetching config",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Success",
			"data":    config,
		})
	}
}

func getCurrentConfigForApp(appName string, dbClient *ent.Client) (*ent.AppConfig, error) {
	config, err := dbClient.AppConfig.
		Query().
		Where(
			appconfig.And(
				appconfig.AppName(appName),
				appconfig.Status("active"),
			),
		).
		Order(appconfig.ByVersion(sql.OrderAsc())).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	if len(config) == 0 {
		return nil, nil
	}

	return config[len(config)-1], nil
}

func CreateConfigEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		if currentUser.Role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
				"error":   nil,
			})
		}

		appName := c.Params("AppName")

		configCreateDto := new(model.CreateConfigDto)

		if err := c.BodyParser(configCreateDto); err != nil {
			log.Println(err)
			return sendBadRequestResponse(c, err, "Invalid request body")
		}

		configVersion := 1

		previousConfig, err := getCurrentConfigForApp(appName, dbClient)

		if err != nil {
			log.Println(err)
			return sendBadRequestResponse(c, err, "Error fetching config")
		}

		status := "active"

		if previousConfig != nil {
			configVersion = previousConfig.Version + 1
			status = "inactive"
		}

		config, err := dbClient.AppConfig.
			Create().
			SetAppName(appName).
			SetVersion(configVersion).
			SetStatus(status).
			SetConfig(configCreateDto.Config).
			Save(context.Background())

		if err != nil {
			log.Println(err)
			return sendBadRequestResponse(c, err, "Error creating config")
		}

		apiAlertsClient := c.Locals("ApiAlertsClient").(*apialerts.Client)
		versionString := strconv.Itoa(configVersion)
		go apiAlertsClient.Send("Config created for app: "+appName,
			[]string{appName, versionString},
			"v1/config/"+appName)

		return c.JSON(fiber.Map{
			"message": "Success",
			"data":    config,
		})
	}
}

func ActivateConfigEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		if currentUser.Role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
				"error":   nil,
			})
		}

		appName := c.Params("AppName")
		version := c.Params("Version")

		if appName == "" {
			return sendBadRequestResponse(c, nil, "AppName is required")
		}

		if version == "" {
			return sendBadRequestResponse(c, nil, "Version is required")
		}

		versionInt, err := strconv.Atoi(version)

		if err != nil {
			return sendBadRequestResponse(c, err, "Invalid version")
		}

		config, err := dbClient.
			AppConfig.
			Query().
			Where(
				appconfig.AppName(appName),
				appconfig.Version(versionInt),
			).First(context.Background())

		if err != nil {
			return sendBadRequestResponse(c, err, "Error fetching config")
		}

		if config.Status != "inactive" {
			return sendBadRequestResponse(c, nil, "Config is already active")
		}

		oldConfig, err := getCurrentConfigForApp(appName, dbClient)

		if err != nil {
			log.Println(err)
			return sendBadRequestResponse(c, err, "Error fetching config")
		}

		if oldConfig != nil {
			oldConfig.Status = "inactive"
			if err := dbClient.AppConfig.UpdateOne(oldConfig).SetStatus("inactive").Exec(context.Background()); err != nil {
				log.Println(err)
				return sendBadRequestResponse(c, err, "Error deactivating old config")
			}
		}

		config.Status = "active"

		if err := dbClient.AppConfig.UpdateOne(config).SetStatus("active").Exec(context.Background()); err != nil {
			log.Println(err)
			return sendBadRequestResponse(c, err, "Error activating config")
		}

		apiAlertsClient := c.Locals("ApiAlertsClient").(*apialerts.Client)
		go apiAlertsClient.Send("Config activated for app: "+appName,
			[]string{appName, version},
			"v1/config/"+appName)

		return c.JSON(fiber.Map{
			"message": "Success",
			"data":    config,
		})
	}
}
