package main

import (
	"log" // 生成されたAPIコードのインポートパスを適切に設定
	"net/http"

	"github.com/labstack/echo/v4"
	api "github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

func main() {
	// Echoインスタンスの初期化
	e := echo.New()

	// 生成されたAPIハンドラーの登録
	api.RegisterHandlers(e, NewServer())

	// サーバの起動
	log.Fatal(e.Start(":8080"))
}

// Server は生成されたAPIインターフェースを実装する
type Server struct{}

// NewServer はServer構造体の新しいインスタンスを作成する
func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetPets(ctx echo.Context) error {

	pets := []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		{ID: 1, Name: "cat"},
		{ID: 2, Name: "dog"},
	}
	// ペットのデータをJSON形式でレスポンスとして返す
	return ctx.JSON(http.StatusOK, pets)
}
func (s *Server) PostPets(ctx echo.Context) error {
	// ペットの追加ロジック
	pet := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID: 1, Name: "cat",
	}
	// ペットのデータをJSON形式でレスポンスとして返す
	return ctx.JSON(http.StatusCreated, pet)
}
func (s *Server) UpdatePetPartial(ctx echo.Context, petId int) error {
	pet := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID: 1, Name: "new cat",
	}
	// ペットのデータをJSON形式でレスポンスとして返す
	return ctx.JSON(http.StatusOK, pet)
}
func (s *Server) DeletePet(ctx echo.Context, petId int) error {
	return ctx.NoContent(http.StatusNoContent)
}
