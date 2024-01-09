//go:generate mockgen -source=service.go -destination=mock_service.go -package=service
package service

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

type TodoAdapter interface {
	Create(ctx context.Context, req CreateTodoRequest) (*api.Todo, error)
	Update(ctx context.Context, req UpdateTodoRequest) (*api.Todo, error)
	List(ctx context.Context, req GetTodosRequest) ([]*api.Todo, error)
	Get(ctx context.Context, req GetTodoRequest) (*api.Todo, error)
	Delete(ctx context.Context, req DeleteTodoRequest) error
}
