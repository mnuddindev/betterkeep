package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mnuddindev/betterkeep/db"
	"github.com/mnuddindev/betterkeep/routes"
	"github.com/mnuddindev/betterkeep/utils"
)

func main() {
	db.Connect()
	port := utils.Config("APP_PORT")
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "http://localhost:5307/",
			AllowCredentials: true,
		}),
		compress.New(compress.Config{
			Level: compress.LevelBestCompression,
		}),
	)
	routes.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	log.Fatal(app.Listen(":" + port))
}
