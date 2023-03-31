package main

import (
	"kurdi-go/infrastructure/database"
	postgresDatabse "kurdi-go/infrastructure/database/postgres"
	"kurdi-go/web/config"
	"kurdi-go/web/middlewares"
	"kurdi-go/web/routes"
	"os"

	"github.com/gofiber/fiber/v2"

	_ "github.com/swaggo/fiber-swagger/example/docs"
)

// @title Inventory API
// @version 1.0
// @description This is an inventory API Docs.
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
	config.ENVConfig()
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
	err := app.Listen(os.Getenv("PORT")); if err != nil{
		return
	}
	//----------------------------

}
