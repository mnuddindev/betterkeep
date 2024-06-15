package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mnuddindev/betterkeep/auth"
	"github.com/mnuddindev/betterkeep/routes"
)

func main() {
	port := auth.Config("APP_PORT")
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	routes.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	log.Fatal(app.Listen(":" + port))
}
