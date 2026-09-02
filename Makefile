# connecting .env file
include .env
export

MAKEFLAGS += --silent		# dont show text at terminal

MIGRATE=migrate

migrate-up: ## Apply all migrations up
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down: ## Apply latest migrations down
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-force: ## Immediatly use version. Example: make migrate-force v=1
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(v)

migrate-version: ## Показать текущую версию миграции
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

migrate-new name: ## Create new migration. Example: make migrate-new name=add_users
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

server-run:
	go run ./cmd/web