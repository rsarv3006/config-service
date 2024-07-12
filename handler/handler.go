package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func sendBadRequestResponse(c *fiber.Ctx, err error, message string) error {
	log.Println(err)
	log.Println(message)
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
		"error":   err,
	})
}

func HealthEndpoint() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	}
}
