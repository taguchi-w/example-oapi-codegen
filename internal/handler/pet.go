package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
)

// Server は生成されたAPIインターフェースを実装する
type Pet struct {
	Pet PetService
}

// NewServer はServer構造体の新しいインスタンスを作成する
func NewPet(pet PetService) *Pet {
	return &Pet{pet}
}

func (h *Pet) GetPets(ctx echo.Context) error {
	pets, err := h.Pet.List(ctx.Request().Context(), service.GetPetsRequest{
		Offset: 0,
		Limit:  20,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, pets)
}

func (h *Pet) PostPets(ctx echo.Context) error {
	pet, err := h.Pet.Create(ctx.Request().Context(), service.CreatePetRequest{
		Pet: api.Pet{
			Id:   "1",
			Name: "cat",
		},
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, pet)
}

func (h *Pet) UpdatePetPartial(ctx echo.Context, petId string) error {
	pet, err := h.Pet.Update(ctx.Request().Context(), service.UpdatePetRequest{
		Id:   1,
		Name: util.P("cat"),
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (h *Pet) DeletePet(ctx echo.Context, petId string) error {
	err := h.Pet.Delete(ctx.Request().Context(), service.DeletePetRequest{
		Id: 1,
	})
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
