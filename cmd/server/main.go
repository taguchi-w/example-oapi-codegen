package main

import (
	"log" // 生成されたAPIコードのインポートパスを適切に設定
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taguchi-w/example-oapi-codegen/internal/adapter"
	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	api "github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
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
type Server struct {
	service.Services
}

// NewServer はServer構造体の新しいインスタンスを作成する
func NewServer() *Server {
	adapters := adapter.New()
	return &Server{Services: service.New(adapters)}
}

func (s *Server) GetPets(ctx echo.Context) error {
	pets, err := s.Pet.List(ctx.Request().Context(), service.GetPetsRequest{
		Offset: 0,
		Limit:  20,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, pets)
}

func (s *Server) PostPets(ctx echo.Context) error {
	pet, err := s.Pet.Create(ctx.Request().Context(), service.CreatePetRequest{
		Pet: api.Pet{
			Id:   1,
			Name: "cat",
		},
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, pet)
}

func (s *Server) UpdatePetPartial(ctx echo.Context, petId int) error {
	pet, err := s.Pet.Update(ctx.Request().Context(), service.UpdatePetRequest{
		Id:   1,
		Name: util.P("cat"),
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, pet)
}
func (s *Server) DeletePet(ctx echo.Context, petId int) error {
	err := s.Pet.Delete(ctx.Request().Context(), service.DeletePetRequest{
		Id: 1,
	})
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
