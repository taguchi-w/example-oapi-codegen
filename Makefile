# Makefile

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
.PHONY: generate
generate:
	oapi-codegen -package $(GENERATED_PACKAGE) -o $(GENERATED_DIR)/$(GENERATED_PACKAGE).gen.go $(SPEC_FILE)

# ビルド
.PHONY: build
build:
	go build -o $(BINARY_NAME) cmd/server/main.go

# クリーンアップ
.PHONY: clean
clean:
	rm $(BINARY_NAME)
	rm $(GENERATED_DIR)/api.gen.go

# 実行
.PHONY: run-server
run-server:
	go run $(SERVER_ENTRY_POINT)
# 実行
.PHONY: run-client
run-client:
	go run $(CLIENT_ENTRY_POINT)

# すべてのタスクを実行
.PHONY: all
all: generate build
