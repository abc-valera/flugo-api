# util commands
enter_db:
	docker exec -it flugo-db bash

# generate commands
init_migrations:
	goose -dir migration/ -s create init_migrations sql

# run commands
run_flugo-db:
	docker run --name flugo-db -p "5432:5432" -e POSTGRES_USER=abc_valera -e POSTGRES_PASSWORD=abc_valera -e POSTGRES_DB=flugo -d postgres
run_migrations-up:
	goose -dir migration postgres "postgresql://abc_valera:abc_valera@localhost:5432/flugo?sslmode=disable" up
run_flugo-api_local:
	(go build -o build/flugo-api cmd/api/main.go;./build/flugo-api)

