
.PHONY: clean test security build run

APP_NAME = inventory
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = infrastructure/database/migrations
DATABASE_URL = postgres://postgres:123456789@localhost/inventory?sslmode=disable

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

# test: security
# 	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
# 	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go
	cp .env build


migration.create:
	migrate create -ext sql -dir $(MIGRATIONS_FOLDER) -seq $(name)

migration.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migration.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migration.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.run: docker.network docker.postgres swag docker.fiber migrate.up

docker.network:
	docker network inspect go-inventory-network >/dev/null 2>&1 || \
	docker network create -d bridge go-inventory-network

docker.fiber.build:
	docker build -t go-inventory .

#change netrowk to be the same of postgres
docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name go-inventory \
		--network maxab_basic_env \ 
		-p 3000:3000 \
		go-inventory

docker.postgres.create:
	docker run --rm -d \
		--name postgres \
		--network go-inventory-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.db.create:
	docker exec -it postgres createdb --username=postgres --owner=postgres inventory

docker.db.drop:
	docker exec -it postgres dropdb --username=postgres inventory

docker.postgres.stop:
	docker stop postgres

docker.fiber.stop:
	docker stop go-inventory

docker.stop: docker.stop.fiber docker.stop.postgres

swag:
	/home/kurdi/go/bin/swag init -o ./web/docs 

mock:
	/home/kurdi/go/bin/mockery --all 

test:
	go test ./test/... 

killport:
	sudo kill -9 `sudo lsof -t -i:3000`
