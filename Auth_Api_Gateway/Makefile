MIGRATION_FOLDER = db/migrations
ADDR = root:3003@tcp(127.0.0.1:3306)/auth_dev

# create new migrations -> gmake migrate-create name="migrationName" 
migrate-create: 
	/home/nitesh/go-workspace/bin/goose -dir $(MIGRATION_FOLDER) create $(name) sql

migrate-up:
	/home/nitesh/go-workspace/bin/goose  -dir $(MIGRATION_FOLDER) mysql "$(ADDR)" up

migrate-down:
	/home/nitesh/go-workspace/bin/goose  -dir $(MIGRATION_FOLDER) mysql "$(ADDR)" down