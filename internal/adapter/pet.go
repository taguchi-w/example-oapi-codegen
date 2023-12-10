package adapter

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

// Pet は生成されたAPIインターフェースを実装する
var _ service.PetManager = (*Pet)(nil)

type Pet struct {
}

func NewPet() *Pet {
	return &Pet{}
}

func (a *Pet) Create(ctx context.Context, req service.CreatePetRequest) (*api.Pet, error) {
	return &api.Pet{
		Id:   1,
		Name: "cat",
	}, nil
}
func (a *Pet) Update(ctx context.Context, req service.UpdatePetRequest) (*api.Pet, error) {
	return &api.Pet{
		Id:   1,
		Name: "new cat",
	}, nil
}
func (a *Pet) List(ctx context.Context, req service.GetPetsRequest) ([]*api.Pet, error) {
	return []*api.Pet{
		{Id: 1, Name: "cat"},
		{Id: 2, Name: "dog"},
	}, nil
}
func (a *Pet) Get(ctx context.Context, req service.GetPetRequest) (*api.Pet, error) {
	return &api.Pet{
		Id:   1,
		Name: "cat",
	}, nil
}
func (a *Pet) Delete(ctx context.Context, req service.DeletePetRequest) error {
	return nil
}
