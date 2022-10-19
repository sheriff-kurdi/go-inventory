package main

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/http/routes"
	database2 "kurdi-go/infrastructure/infrastructure_database"
)

func main() {
	//app
	app := fiber.New()

	//Database
	// Connect to database
	database2.Connect()
	//migrate database
	//TODO:Migrate based on env variable
	database2.PostgresAutoMigrate()

	// Routes
	routes.StockRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}

}
