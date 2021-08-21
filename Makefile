DB_URL_ENV=DATABASE_URL
DB_MIGRATIONS_DIR="./migrations"
DB_SCHEMA_FILE="./internal/db/schema.sql"

dev:
	air

migrate-new:
	dbmate -e $(DB_URL_ENV) -d $(DB_MIGRATIONS_DIR) -s $(DB_SCHEMA_FILE) new $(NAME)

migrate-up:
	dbmate -e $(DB_URL_ENV) -d $(DB_MIGRATIONS_DIR) -s $(DB_SCHEMA_FILE) up

migrate-down:
	dbmate -e $(DB_URL_ENV) -d $(DB_MIGRATIONS_DIR) -s $(DB_SCHEMA_FILE) down

migrate-status:
	dbmate -e $(DB_URL_ENV) -d $(DB_MIGRATIONS_DIR) -s $(DB_SCHEMA_FILE) status

sqlc-generate:
	sqlc generate

