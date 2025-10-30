package router

import (
	"myapp/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users", handler.CreateUser)
}
