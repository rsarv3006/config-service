package main

import (
	"RjsConfigService/alert"
	"RjsConfigService/config"
	"RjsConfigService/database"
	"RjsConfigService/router"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func initApp() *fiber.App {
	jwtSecret := config.Config("JWT_SECRET")
	env := config.Config("ENV")

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
	})

	client := database.Connect()
	go database.CreateUserAccounts(client, jwtSecret)

	if env != "production" {
		log.Println("Enabling pprof...")
		app.Use(pprof.New())
	}

	app.Use(helmet.New())
	app.Use(recover.New())

	apiAlertsClient := alert.Connect()

	log.Println("Setting context")
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Locals("JwtSecret", jwtSecret)
		c.Locals("Env", env)
		c.Locals("ApiAlertsClient", apiAlertsClient)
		return c.Next()
	})

	log.Println("Created new fiber app...")
	router.SetupRoutes(app, client)
	log.Println("Routes setup.")

	return app
}

func main() {
	app := initApp()

	log.Println("Listening on port 3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
