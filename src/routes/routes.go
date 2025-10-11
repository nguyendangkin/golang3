package routes

import (
	"natasha/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/register", users.RegisterUser)
}
