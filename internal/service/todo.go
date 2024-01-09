package service

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

type Todo struct {
	todoAdapter TodoAdapter
}

func NewTodo(todoAdapter TodoAdapter) *Todo {
	return &Todo{todoAdapter}
}

type CreateTodoRequest struct {
	api.Todo
}

func (r CreateTodoRequest) Validate() error {
	return nil
}

type GetTodosRequest struct {
	Offset int
	Limit  int
}

func (r GetTodosRequest) Validate() error {
	return nil
}

type GetTodoRequest struct {
	Id string
}

func (r GetTodoRequest) Validate() error {
	return nil
}

type UpdateTodoRequest struct {
	Id      string
	Subject *string
	Body    *string
}

func (r UpdateTodoRequest) Validate() error {
	return nil
}

type DeleteTodoRequest struct {
	Id string
}

func (r DeleteTodoRequest) Validate() error {
	return nil
}

func (s *Todo) Create(ctx context.Context, req CreateTodoRequest) (*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pet, err := s.todoAdapter.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return pet, nil
}
func (s *Todo) List(ctx context.Context, req GetTodosRequest) ([]*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pets, err := s.todoAdapter.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return pets, nil
}
func (s *Todo) Get(ctx context.Context, req GetTodoRequest) (*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pet, err := s.todoAdapter.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return pet, nil
}
func (s *Todo) Update(ctx context.Context, req UpdateTodoRequest) (*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pet, err := s.todoAdapter.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return pet, nil
}
func (s *Todo) Delete(ctx context.Context, req DeleteTodoRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	err := s.todoAdapter.Delete(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
