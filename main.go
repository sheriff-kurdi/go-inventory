package main

import (
	"kurdi-go/infrastructure/database"
	postgresDatabse "kurdi-go/infrastructure/database/postgres"
	"kurdi-go/web/middlewares"
	"kurdi-go/web/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/swaggo/fiber-swagger/example/docs"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
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
	middlewares.FiberMiddleware(app) // Register Fiber's middleware for app.
	routes.SwaggerRoute(app)
	routes.ProductsRoutes(app)
	//----------------------------



	//---------Port-------------
	err = app.Listen(":3000")
	if err != nil {
		return
	}
	//----------------------------

}
