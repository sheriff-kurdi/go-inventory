package main

import (
	"kurdi-go/infrastructure/database"
	postgresDatabse "kurdi-go/infrastructure/database/postgres"
	"kurdi-go/web/routes"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	//app
	app := fiber.New()

	//----------env file loading-------
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//----------------------------

	//----------Database-------
	connection := postgresDatabse.Connect()
	database.AutoMigrate(connection)
	//----------------------------

	//---------Routes-------------
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	routes.ProductsRoutes(app)
	//----------------------------

	//---------Port-------------
	err = app.Listen(":3000")
	if err != nil {
		return
	}
	//----------------------------

}
