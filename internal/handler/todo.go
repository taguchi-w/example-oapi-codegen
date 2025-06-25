package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
)

// Server は生成されたAPIインターフェースを実装する
type Todo struct {
	Todo TodoService
}

// NewServer はServer構造体の新しいインスタンスを作成する
func NewTodo(pet TodoService) *Todo {
	return &Todo{pet}
}

func (h *Todo) GetTodos(ctx echo.Context) error {
	pets, err := h.Todo.List(ctx.Request().Context(), service.GetTodosRequest{
		Offset: 0,
		Limit:  20,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, pets)
}

func (h *Todo) PostTodos(ctx echo.Context) error {
	var req api.Todo
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pet, err := h.Todo.Create(ctx.Request().Context(), service.CreateTodoRequest{
		Todo: api.Todo{
			Subject: req.Subject,
			Body:    req.Body,
		},
	})
	if err != nil {

		return err
	}
	return ctx.JSON(http.StatusCreated, pet)
}

func (h *Todo) UpdateTodoPartial(ctx echo.Context, todoId string) error {
	var req api.Todo
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pet, err := h.Todo.Update(ctx.Request().Context(), service.UpdateTodoRequest{
		Id:      todoId,
		Subject: util.PorNil(req.Subject, ""),
		Body:    util.PorNil(req.Body, ""),
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (h *Todo) DeleteTodo(ctx echo.Context, todoId string) error {
	err := h.Todo.Delete(ctx.Request().Context(), service.DeleteTodoRequest{
		Id: todoId,
	})
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
