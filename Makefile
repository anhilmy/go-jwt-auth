CONFIG_FILE ?= .local.yml
DB_USER ?= $(shell sed -n 's/^db_user:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
DB_PASSWORD ?= $(shell sed -n 's/^db_password:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
DB_NAME ?= $(shell sed -n 's/^db_name:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
DB_PORT ?= $(shell sed -n 's/^db_port:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
DB_HOST ?= $(shell sed -n 's/^db_host:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
APP_DSN = postgres://$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable&user=$(DB_USER)&password=$(DB_PASSWORD)
MIGRATE := docker run --rm -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.10.0 -path=/migrations/ -database "$(APP_DSN)"

.PHONY: db-init
db-init: ## start the database server
	@mkdir -m 777 -p testdata/postgres
	docker run --rm --name "$(DB_NAME)" -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_USER="$(DB_USER)" -e POSTGRES_PASSWORD="$(DB_PASSWORD)" -e POSTGRES_DB="$(DB_NAME)" -d -p $(DB_PORT):5432 postgis/postgis:14-3.2 \
		-c hba_file=/testdata/postgresql.conf
	docker exec $(DB_NAME) createdb -U postgres $(DB_NAME)

.PHONY: db-start
db-start:
	docker run --rm --name "$(DB_NAME)" -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_USER="$(DB_USER)" -e POSTGRES_PASSWORD="$(DB_PASSWORD)" -e POSTGRES_DB="$(DB_NAME)" -d -p $(DB_PORT):5432 postgis/postgis:14-3.2 \
		-c hba_file=/testdata/postgresql.conf

.PHONY: db-stop
db-stop: 
	docker stop "$(DB_NAME)"

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	migrate create -ext sql -dir $(shell pwd)/migrations -seq $${name}

.PHONY: echo
echo:
	@echo $(MIGRATE)