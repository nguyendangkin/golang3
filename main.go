package main

import (
	"fmt"
	"log"
	"natasha/src/config"
	"natasha/src/database"
	"natasha/src/modules/users"
	"natasha/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	database.Connect()

	// auto migrate
	if err := database.Repo.AutoMigrate(&users.User{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Auto migrated")

	app := fiber.New()
	routes.RegisterRoutes(app)
	app.Listen(fmt.Sprintf(":%d", config.Data.Server.Port))

}
