package sqlserver

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect() *gorm.DB {
	if connection != nil {
		return connection
	}
	dsn := os.Getenv("SQL_SERVER_DATABASE_CONNECTION")
	postgresDb, postgresErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if postgresErr != nil {
		panic(postgresErr)
	}
	connection = postgresDb
	return connection

}
