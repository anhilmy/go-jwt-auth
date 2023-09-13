CONFIG_FILE ?= .local.yml
APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
MIGRATE := docker run --rm -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.10.0 -path=/migrations/ -database "$(APP_DSN)"
DB_USER ?= $(shell sed -n 's/^db_user:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
DB_PASSWORD ?= $(shell sed -n 's/^db_password:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
DB_NAME ?= $(shell sed -n 's/^db_name:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))

.PHONY: db-start
db-start: ## start the database server
	@mkdir -p testdata/postgres
	docker run --rm --name postgres -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGERS_USER="$(DB_USER)" -e POSTGRES_PASSWORD="$(DB_PASSWORD)" -e POSTGRES_DB="$(DB_NAME)" -d -p 5432:5432 postgis/postgis:14-3.2
		--platform linux/arm64/v8