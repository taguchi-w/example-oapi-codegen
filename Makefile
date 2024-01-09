# Makefile

# ------------
# go build
# ------------

# ビルド出力先
BINARY_NAME=bin/server

# OpenAPI仕様ファイル
SPEC_FILE=docs/api/api.yaml

# 生成されるGoコードのパッケージと出力先
GENERATED_DIR=pkg/api
GENERATED_PACKAGE=api
# エントリーポイント
SERVER_ENTRY_POINT=cmd/server/main.go
CLIENT_ENTRY_POINT=cmd/client/main.go

# コード生成
.PHONY: generate-api
generate-api:
	oapi-codegen -package $(GENERATED_PACKAGE) -o $(GENERATED_DIR)/$(GENERATED_PACKAGE).gen.go $(SPEC_FILE)

.PHONY: generate-mock
generate-mock:
	go generate ./...

.PHONY: generate
generate:
	make generate-api
	make generate-mock

.PHONY: lint
lint:
	staticcheck ./...	

.PHONY: test
test:
	go test ./...	

.PHONY: build
build:
	go build -o $(BINARY_NAME) cmd/server/main.g

.PHONY: clean
clean:
	rm $(BINARY_NAME)
	rm $(GENERATED_DIR)/api.gen.go

.PHONY: run-server
run-server:
	go run $(SERVER_ENTRY_POINT)

.PHONY: run-client
run-client:
	go run $(CLIENT_ENTRY_POINT)

# ------------
# sql-migrate
# ------------
SQL_MIGRATE_DIR=./assets/sql-migrate 
SQL_MIGRATE_ENV=dev
SQL_MIGRATE_DB=oapicodegen
SQL_DUMP_DIR=./assets/docker/mysql/ddl

.PHONY: sql-migrate-state
sql-migrate-state:
	cd $(SQL_MIGRATE_DIR)/ && sql-migrate status -env=$(SQL_MIGRATE_ENV)
.PHONY: sql-migrate-up
sql-migrate-up:
	cd $(SQL_MIGRATE_DIR)/ && sql-migrate up -env=$(SQL_MIGRATE_ENV)
.PHONY: sql-migrate-down
sql-migrate-down:
	cd $(SQL_MIGRATE_DIR)/ && sql-migrate down -env=$(SQL_MIGRATE_ENV)
.PHONY: sql-dump
sql-dump:
	echo "USE $(SQL_MIGRATE_DB);" > $(SQL_DUMP_DIR)/003_create_tables.sql
	mysqldump -u root --protocol=tcp --no-data $(SQL_MIGRATE_DB) >> $(SQL_DUMP_DIR)/003_create_tables.sql
	echo "USE $(SQL_MIGRATE_DB);" > $(SQL_DUMP_DIR)/004_init_tabels.sql
	mysqldump -u root --protocol=tcp --no-create-info $(SQL_MIGRATE_DB) gorp_migrations >> $(SQL_DUMP_DIR)/004_init_tabels.sql

# すべてのタスクを実行
.PHONY: all
all: generate lint test build
