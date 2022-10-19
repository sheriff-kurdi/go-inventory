startPostgres:
	sudo docker container start postgres

createDB:
	sudo docker exec -it postgres createdb --username=postgres --owner=postgres inventory

dropDB:
	sudo docker exec -it postgres dropdb --username=postgres inventory

migrateUp:
	migrate -path database/migrations -database "postgresql://postgres:123456789@localhost:5432/inventory?sslmode=disable" -verbose up

migrateDown:
	migrate -path database/migrations -database "postgresql://postgres:123456789@localhost:5432/inventory?sslmode=disable" -verbose down

createMigration:
	migrate create -ext sql -dir infrastructure/infrastructure_database/migrations -seq $(name)

buildUp:	startPostgres createDB migrateUp

