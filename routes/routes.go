package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnuddindev/betterkeep/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/users")
	auth := api.Group("/auth")

	// Home
	api.Get("/home", controllers.NotImplemented)

	// user system
	user.Get("/", controllers.NotImplemented)
	user.Get("/:id", controllers.NotImplemented)
	user.Post("/:id", controllers.NotImplemented)
	user.Put("/:id", controllers.NotImplemented)
	user.Delete("/:id", controllers.NotImplemented)

	// login system
	auth.Post("/register", controllers.CreateUser)
	auth.Post("/login", controllers.NotImplemented)
	auth.Post("/forget-password", controllers.NotImplemented)
	auth.Get("/refresh", controllers.NotImplemented)
	auth.Get("/active-user/:userid", controllers.ActiveUser)
}
