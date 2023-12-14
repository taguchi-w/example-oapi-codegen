package adapter

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

type Pet struct {
	db DBAdapter
	id IDGenerator
}

func NewPet(db DBAdapter, id IDGenerator) *Pet {
	return &Pet{db, id}
}

func (a *Pet) Create(ctx context.Context, req service.CreatePetRequest) (*api.Pet, error) {
	req.Pet.Id = a.id.Generate()
	_, err := a.db.NamedExec(`INSERT INTO pets (id, name, tag) VALUES (:id,:name,:tag)`, map[string]interface{}{
		"id":   req.Pet.Id,
		"name": req.Pet.Name,
		"tag":  req.Pet.Tag,
	})
	if err != nil {
		return nil, err
	}
	return &req.Pet, nil
}
func (a *Pet) Update(ctx context.Context, req service.UpdatePetRequest) (*api.Pet, error) {
	return &api.Pet{
		Id:   "1",
		Name: "new cat",
	}, nil
}
func (a *Pet) List(ctx context.Context, req service.GetPetsRequest) ([]*api.Pet, error) {
	return []*api.Pet{
		{Id: "1", Name: "cat"},
		{Id: "2", Name: "dog"},
	}, nil
}
func (a *Pet) Get(ctx context.Context, req service.GetPetRequest) (*api.Pet, error) {
	return &api.Pet{
		Id:   "1",
		Name: "cat",
	}, nil
}
func (a *Pet) Delete(ctx context.Context, req service.DeletePetRequest) error {
	return nil
}
