package main

import (
	"RjsConfigService/router"
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")
	env := os.Getenv("ENV")

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
	})

	println("Initializing emailer...")

	if env != "production" {
		println("Enabling pprof...")
		app.Use(pprof.New())
	}

	app.Use(helmet.New())
	app.Use(recover.New())

	println("Setting context")
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		c.Locals("JwtSecret", jwtSecret)
		c.Locals("Env", env)
		c.Locals("APPLE_CODE", os.Getenv("APPLE_CODE"))

		return c.Next()
	})

	println("Created new fiber app...")

	router.SetupRoutes(app)

	println("Routes setup.")

	err := app.Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}

}
