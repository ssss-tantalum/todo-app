NAME :=

build:
	@templ generate
	@npm run build
	@go build -o ./bin/app ./cmd/app/main.go

run: generate build
	@./bin/app

dev:
	@air

init:
	@templ generate && npm install

generate:
	@templ generate && npm run build

db-init:
	@go run cmd/migrate/main.go db init

db-migrate:
	@go run cmd/migrate/main.go db migrate

db-rollback:
	@go run cmd/migrate/main.go db rollback

db-create_go:
	$(call __require_NAME)
	@go run cmd/migrate/main.go db create_go ${NAME}

db-create_sql:
	$(call __require_NAME)
	@go run cmd/migrate/main.go db create_sql ${NAME}

db-lock:
	@go run cmd/migrate/main.go db lock

db-unlock:
	@go run cmd/migrate/main.go db unlock

db-status:
	@go run cmd/migrate/main.go db status

db-mark_applied:
	@go run cmd/migrate/main.go db mark_applied

define __require_NAME
    @bash -c "if [ '${NAME}' = '' ]; then echo 'NAME is not defined; you must specify NAME like $$ make NAME=create_xxx_table task'; exit 1; fi"
endef