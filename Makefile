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


## 依存関係のインストール
.PHONY: deps
deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/golang/mock/mockgen@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
ifneq ($(shell command -v asdf 2> /dev/null),)
	asdf reshim
endif

# ------------------------
# generate code
# ------------------------

## OpenAPI仕様ファイルからコード生成
.PHONY: generate-api
generate-api:
	oapi-codegen -package $(GENERATED_PACKAGE) -o $(GENERATED_DIR)/$(GENERATED_PACKAGE).gen.go $(SPEC_FILE)

## mockの生成
.PHONY: generate-mock
generate-mock:
	go generate ./...

## generateされたコードを削除
.PHONY: generate-clean
clean:
	rm $(BINARY_NAME)
	rm $(GENERATED_DIR)/api.gen.go

## コード生成
.PHONY: generate
generate:
	make generate-clean
	make generate-api
	make generate-mock

# ------------------------
# test
# ------------------------

## Lintチェック
.PHONY: lint
lint:
	staticcheck ./...	

## 単体テスト
.PHONY: test
test:
	go test ./...	

## Docker Composeを使用してMySQLコンテナを起動し、テストを実行
.PHONY: integration-test
integration-test:
	docker compose -f docker-compose.yml run --build --rm -e MYSQL_DSN="root@tcp(mysql:3306)/oapicodegen?parseTime=true" test


.PHONY: build
build:
	go build -o $(BINARY_NAME) cmd/server/main.go


## サーバーの実行
.PHONY: run-server
run-server:
	MYSQL_DSN="root@tcp(localhost:3306)/oapicodegen?parseTime=true" go run $(SERVER_ENTRY_POINT)

## テストクライアントの実行
.PHONY: run-client
run-client:
	go run $(CLIENT_ENTRY_POINT)

# ------------------------
# sql migrate
# ------------------------
SQL_MIGRATE_DIR=./assets/sql-migrate 
SQL_MIGRATE_ENV=dev
SQL_MIGRATE_DB=oapicodegen
SQL_DUMP_DIR=./assets/docker/mysql/ddl

## DBマイクレーション確認
.PHONY: sql-migrate-state
sql-migrate-state:
	cd $(SQL_MIGRATE_DIR)/ && sql-migrate status -env=$(SQL_MIGRATE_ENV)

## DBマイクレーション実行
.PHONY: sql-migrate-up
sql-migrate-up:
	cd $(SQL_MIGRATE_DIR)/ && sql-migrate up -env=$(SQL_MIGRATE_ENV)

## DBマイクレーションロールバック
.PHONY: sql-migrate-down
sql-migrate-down:
	cd $(SQL_MIGRATE_DIR)/ && sql-migrate down -env=$(SQL_MIGRATE_ENV)

## マイクレーション用DDL作成
.PHONY: sql-dump
sql-dump:
	echo "USE $(SQL_MIGRATE_DB);" > $(SQL_DUMP_DIR)/003_create_tables.sql
	mysqldump -u root --protocol=tcp --no-data $(SQL_MIGRATE_DB) >> $(SQL_DUMP_DIR)/003_create_tables.sql
	echo "USE $(SQL_MIGRATE_DB);" > $(SQL_DUMP_DIR)/004_init_tabels.sql
	mysqldump -u root --protocol=tcp --no-create-info $(SQL_MIGRATE_DB) gorp_migrations >> $(SQL_DUMP_DIR)/004_init_tabels.sql


.PHONY: help
help:
	make2help

## すべてのタスクを実行
.DEFAULT_GOAL := default
default: deps generate lint test build