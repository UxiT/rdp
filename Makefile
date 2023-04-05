DB_USER=root
DB_PASSWORD=secret
DB_NAME=lab
MIGRATIONS_DIR=./db/migrations
POSTGRES_URL=postgres://root:secret@127.0.0.1:5432/lab?sslmode=disable

.PHONY: create-migration
create-migration:
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq ${NAME}

.PHONY: run-migrations
run-migrations:
	@migrate -database ${POSTGRES_URL} -path ${MIGRATIONS_DIR} up

.PHONY: rollback-migration
rollback-migration:
	@migrate -database ${POSTGRES_URL} -path $(MIGRATIONS_DIR) down 1
