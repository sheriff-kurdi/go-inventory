package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"kurdi-go/http/routes"
	database2 "kurdi-go/infrastructure/infrastructure_database"
	"log"
)

func main() {
	//app
	app := fiber.New()

	//.env access
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Database
	// Connect to database
	database2.Connect()
	//migrate database
	//TODO:Migrate based on env variable
	database2.PostgresAutoMigrate()

	// Routes
	routes.StockRoutes(app)
	err = app.Listen(":3000")
	if err != nil {
		return
	}

}
