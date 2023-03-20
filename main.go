package main

import (
	"kurdi-go/infrastructure/database"
	postgresDatabse "kurdi-go/infrastructure/database/postgres"
	"kurdi-go/web/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/swaggo/fiber-swagger/example/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	//--------swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	// docs.SwaggerInfo.Title = "Swagger Example API"
	// docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	// docs.SwaggerInfo.BasePath = "/v2"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//----------end swagger--------

	//---------Port-------------
	err = app.Listen(":3000")
	if err != nil {
		return
	}
	//----------------------------

}
