//go:generate mockgen -source=handler.go -destination=mock_handler.go -package=handler
package handler

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

type TodoService interface {
	Create(ctx context.Context, req service.CreateTodoRequest) (*api.Todo, error)
	List(ctx context.Context, req service.GetTodosRequest) ([]*api.Todo, error)
	Get(ctx context.Context, req service.GetTodoRequest) (*api.Todo, error)
	Update(ctx context.Context, req service.UpdateTodoRequest) (*api.Todo, error)
	Delete(ctx context.Context, req service.DeleteTodoRequest) error
}
