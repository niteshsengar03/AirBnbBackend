MIGRATIONS_DIR = db/migrations
DB_URL = root:3003@tcp(127.0.0.1:3306)/auth_dev

# create new migrations -> gmake migrate-create name="migrationName" 
migrate-create: 
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATIONS_DIR) create $(name) sql

migrate-up:
	/home/nitesh/go-workspace/bin/goose  -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up

migrate-down:
	/home/nitesh/go-workspace/bin/goose  -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down

# Rollback all migrations and reset database # gmake migrate-reset
migrate-reset:
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" reset

# Show current migration status # gmake migrate-status
migrate-status:
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) # gmake migrate-redo
migrate-redo:
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" redo

# Run specific migration version # gmake migrate-version version=20200101120000
migrate-to:
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version # gmake migrate-down-to version=20200101120000
migrate-down-to:
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version # gmake migrate-force version=20200101120000
migrate-force:
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" force $(version)

# Print Goose help # gmake migrate-help
migrate-help:
	/home/nitesh/go-workspace/bin/goose -h