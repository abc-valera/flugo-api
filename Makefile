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

# local run commands
run_test_repository:
	go test -count=1 ./internal/infrastructure/repository/...
run_test_service:
	go test -count=1 ./internal/service/...
run_test_delivery:
	go test -count=1 ./internal/delivery/...
run_test_all:
	go test -count=1 ./...
run_flugo-db:
	docker run --name flugo-db -p "5432:5432" -e POSTGRES_USER=abc_valera -e POSTGRES_PASSWORD=abc_valera -e POSTGRES_DB=flugo -d postgres:15-alpine
run_flugo-api_local:
	go build -o build/flugo-api cmd/main.go
	./build/flugo-api
run_all_local:
	docker rm -f flugo-db
	make run_flugo-db
	sleep 2
	make run_flugo-api_local

	