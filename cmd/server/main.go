package main

import (
	"myapp/internal/config"
	"myapp/internal/database"
	"myapp/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.ConnectMongo()

	app := fiber.New()
	router.SetupRoutes(app)

	port := config.GetEnv("PORT")
	app.Listen(":" + port)
}
