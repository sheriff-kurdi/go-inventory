package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	routes.StockRoutes(app)
	err = app.Listen(":3000")
	if err != nil {
		return
	}

}
