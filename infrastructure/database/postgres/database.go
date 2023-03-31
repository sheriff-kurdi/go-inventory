package postgres

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connection *gorm.DB

func Connect() *gorm.DB {
	if connection != nil {
		return connection
	}
	dsn := os.Getenv("POSTGRES_DATABASE_CONNECTION")
	postgresDb, postgresErr := gorm.Open(postgres.Open(dsn), 
	&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if postgresErr != nil {
		panic(postgresErr)
	}
	connection = postgresDb
	return connection

}
