package handler

import "github.com/gofiber/fiber/v2"

func FetchConfigEndpoint(c *fiber.Ctx) error {
	appName := c.Params("AppName")
	return c.JSON(fiber.Map{
		"message": "Success",
		"data": map[string]string{
			"appName": appName,
		},
	})
}
