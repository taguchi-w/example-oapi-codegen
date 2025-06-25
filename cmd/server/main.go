package main

import (
	"log" // 生成されたAPIコードのインポートパスを適切に設定
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/taguchi-w/example-oapi-codegen/internal/adapter"
	"github.com/taguchi-w/example-oapi-codegen/internal/handler"
	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	api "github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

func main() {
	// Echoインスタンスの初期化
	e := echo.New()
	e.Debug = true

	// DB接続
	dsn := os.Getenv("MYSQL_DSN")
	if len(dsn) == 0 {
		log.Fatalf("MYSQL_DSN environment variable is not set")
	}
	db, err := sqlx.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// 生成されたAPIハンドラーの登録
	adapters := NewAdapters(db)
	services := NewServices(adapters)
	handlers := NewHandlers(services)
	api.RegisterHandlers(e, handlers)

	// サーバの起動
	log.Fatal(e.Start(":8080"))
}

// dependency injection

type Handlers struct {
	*handler.Todo
}
type Services struct {
	Todo *service.Todo
}
type Adapters struct {
	Todo *adapter.Todo
}

func NewHandlers(services Services) Handlers {
	return Handlers{
		Todo: handler.NewTodo(services.Todo),
	}
}
func NewServices(adapters Adapters) Services {
	return Services{
		Todo: service.NewTodo(adapters.Todo),
	}
}
func NewAdapters(db adapter.DBAdapter) Adapters {
	return Adapters{
		Todo: adapter.NewTodo(db, adapter.NewXIDGenerator()),
	}
}

type MyStruct struct {
	Field1 string
	Field2 int
}
