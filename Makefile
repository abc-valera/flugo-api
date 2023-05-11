# util commands
enter_db:
	docker exec -it flugo-db bash

# generate commands
init_swagger:
	swag fmt -d internal/delivery
	swag init --parseDependency -o ./docs/swagger -d internal/delivery/handler -g ../api/app.go
init_dbdocs:
	dbdocs build ./docs/dbdocs/db.dbml
init_migrations:
	migrate create -ext sql -dir migration -seq schema
generate_mock:
	mockery --dir internal/infrastructure/pkg --output internal/infrastructure/pkg/mock --all
	mockery --dir internal/infrastructure/repository --output internal/infrastructure/repository/mock --all

# test commands
run_test_repository:
	go test -count=1 ./internal/infrastructure/repository/...
run_test_service:
	go test -count=1 ./internal/service/...
run_test_delivery:
	go test -count=1 ./internal/delivery/...
run_test_all:
	go test -count=1 ./...

# Docker commands
build_flugo-api_image:
	docker build -t flugo-api:latest .

# docker container run commands
run_flugo-db_container:
	docker run --name flugo-db --network flugo-net -p "5432:5432" -e POSTGRES_USER=abc_valera -e POSTGRES_PASSWORD=abc_valera -e POSTGRES_DB=flugo -d postgres:15-alpine
run_flugo-api_container:
	docker run --rm --name flugo --network flugo-net -p 3000:3000 -e DATABASE_URL="postgresql://abc_valera:abc_valera@flugo-db:5432/flugo?sslmode=disable" flugo:latest

# local run commands
run_flugo-api_local:
	go build -o build/flugo-api cmd/main.go
	./build/flugo-api
run_all_local:
	docker rm -f flugo-db
	make run_flugo-db_container
	sleep 2
	make run_flugo-api_local
