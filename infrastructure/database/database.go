package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_CONNECTION")
	postgresDb, postgresErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if postgresErr != nil {
		panic(postgresErr)
	}
	PostgresDB = postgresDb

}
