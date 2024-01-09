package adapter

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
)

type Todo struct {
	db DBAdapter
	id IDGenerator
}

func NewTodo(db DBAdapter, id IDGenerator) *Todo {
	return &Todo{db, id}
}

func (a *Todo) Create(ctx context.Context, req service.CreateTodoRequest) (*api.Todo, error) {
	req.Todo.Id = a.id.Generate()
	_, err := a.db.NamedExec(`INSERT INTO pets (id, name, tag) VALUES (:id,:name,:tag)`, map[string]interface{}{
		"id":      req.Todo.Id,
		"subject": req.Todo.Subject,
		"body":    req.Todo.Body,
	})
	if err != nil {
		return nil, err
	}
	return &req.Todo, nil
}
func (a *Todo) Update(ctx context.Context, req service.UpdateTodoRequest) (*api.Todo, error) {
	return &api.Todo{
		Id:      "1",
		Subject: "subject a",
	}, nil
}
func (a *Todo) List(ctx context.Context, req service.GetTodosRequest) ([]*api.Todo, error) {
	return []*api.Todo{
		{Id: "1", Subject: "subject a", Body: util.P("body")},
		{Id: "2", Subject: "subject b", Body: util.P("body")},
	}, nil
}
func (a *Todo) Get(ctx context.Context, req service.GetTodoRequest) (*api.Todo, error) {
	return &api.Todo{
		Id:      "1",
		Subject: "subject a",
	}, nil
}
func (a *Todo) Delete(ctx context.Context, req service.DeleteTodoRequest) error {
	return nil
}
