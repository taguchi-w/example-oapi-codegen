package main

import (
	"log" // 生成されたAPIコードのインポートパスを適切に設定

	"github.com/labstack/echo/v4"
	"github.com/taguchi-w/example-oapi-codegen/internal/adapter"
	"github.com/taguchi-w/example-oapi-codegen/internal/handler"
	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	api "github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

func main() {
	// Echoインスタンスの初期化
	e := echo.New()

	// 生成されたAPIハンドラーの登録
	db := interface{}(nil) // 本来はDBのインスタンスを渡す

	adapters := NewAdapters(db)
	// services := service.New(adapters)
	services := NewServices(adapters)
	handlers := NewHandlers(services)
	api.RegisterHandlers(e, handlers)

	// サーバの起動
	log.Fatal(e.Start(":8080"))
}

// dependency injection

type Handlers struct {
	*handler.Pet
}
type Services struct {
	Pet *service.Pet
}
type Adapters struct {
	Pet *adapter.Pet
}

func NewHandlers(services Services) Handlers {
	return Handlers{
		Pet: handler.NewPet(services.Pet),
	}
}
func NewServices(adapters Adapters) Services {
	return Services{
		Pet: service.NewPet(adapters.Pet),
	}
}
func NewAdapters(db interface{}) Adapters {
	return Adapters{
		Pet: adapter.NewPet(),
	}
}

type MyStruct struct {
	Field1 string
	Field2 int
}
