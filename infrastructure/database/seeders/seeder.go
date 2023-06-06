package seeders

import "github.com/sheriff-kurdi/inventory/infrastructure/database/postgres"

func Seed(){
	connection := postgres.Connect()
	ProductsSeed(connection)
}