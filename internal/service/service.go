//go:generate mockgen -source=service.go -destination=mock_service.go -package=service
package service

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

type PetAdapter interface {
	Create(ctx context.Context, req CreatePetRequest) (*api.Pet, error)
	Update(ctx context.Context, req UpdatePetRequest) (*api.Pet, error)
	List(ctx context.Context, req GetPetsRequest) ([]*api.Pet, error)
	Get(ctx context.Context, req GetPetRequest) (*api.Pet, error)
	Delete(ctx context.Context, req DeletePetRequest) error
}
