//go:generate mockgen -source=handler.go -destination=mocks/mock_handler.go -package=mock
package handler

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

type PetService interface {
	Create(ctx context.Context, req service.CreatePetRequest) (*api.Pet, error)
	List(ctx context.Context, req service.GetPetsRequest) ([]*api.Pet, error)
	Get(ctx context.Context, req service.GetPetRequest) (*api.Pet, error)
	Update(ctx context.Context, req service.UpdatePetRequest) (*api.Pet, error)
	Delete(ctx context.Context, req service.DeletePetRequest) error
}
