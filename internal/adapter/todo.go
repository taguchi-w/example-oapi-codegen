package adapter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

var TodoColumns = []string{
	"id",
	"subject",
	"body",
	"created_at",
	"updated_at",
}

type TodoValue struct {
	Id        string    `db:"id"`
	Subject   string    `db:"subject"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (v *TodoValue) ToModel() *api.Todo {
	if v == nil {
		return nil
	}
	return &api.Todo{
		Id:        v.Id,
		Subject:   v.Subject,
		Body:      v.Body,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

type Todo struct {
	db DBAdapter
	id IDGenerator
}

func NewTodo(db DBAdapter, id IDGenerator) *Todo {
	return &Todo{db, id}
}

func (a *Todo) Create(ctx context.Context, req service.CreateTodoRequest) (*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	req.Todo.Id = a.id.Generate()
	_, err := a.db.NamedExec(`INSERT INTO todo (id, subject, body) VALUES (:id,:subject,:body)`, map[string]interface{}{
		"id":      req.Todo.Id,
		"subject": req.Todo.Subject,
		"body":    req.Todo.Body,
	})
	if err != nil {
		return nil, err
	}
	todo, err := a.Get(ctx, service.GetTodoRequest{Id: req.Todo.Id})
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (a *Todo) Update(ctx context.Context, req service.UpdateTodoRequest) (*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	values := map[string]interface{}{
		"id": req.Id,
	}
	sets := []string{
		"updated_at = NOW()",
	}
	if req.Subject != nil {
		values["subject"] = *req.Subject
		sets = append(sets, "subject = :subject")
	}
	if req.Body != nil {
		values["body"] = *req.Body
		sets = append(sets, "body = :body")
	}

	res, err := a.db.NamedExec(
		fmt.Sprintf(`UPDATE todo SET %s WHERE id = :id LIMIT 1`, strings.Join(sets, ",")),
		values,
	)
	if err != nil {
		return nil, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		fmt.Println("Error updating todo:", values)
		return nil, service.ErrNotFound
	}
	todo, err := a.Get(ctx, service.GetTodoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (a *Todo) List(ctx context.Context, req service.GetTodosRequest) ([]*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	values := map[string]interface{}{
		"offset": req.Offset,
		"limit":  req.Limit,
	}
	// var rows []*TodoValue
	rows, err := a.db.NamedQuery(
		fmt.Sprintf("SELECT %s FROM todo LIMIT :offset,:limit", strings.Join(TodoColumns, ",")),
		values,
	)
	if err != nil {
		return nil, err
	}
	var todos []*api.Todo
	for rows.Next() {
		var v TodoValue
		if err := rows.StructScan(&v); err != nil {
			return nil, err
		}
		todos = append(todos, v.ToModel())
	}
	return todos, nil
}
func (a *Todo) Get(ctx context.Context, req service.GetTodoRequest) (*api.Todo, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	var row TodoValue
	err := a.db.Get(&row,
		fmt.Sprintf(
			"SELECT %s FROM todo WHERE id = ? LIMIT 1",
			strings.Join(TodoColumns, ","),
		),
		req.Id,
	)
	if err != nil {
		return nil, err
	}
	return row.ToModel(), nil
}
func (a *Todo) Delete(ctx context.Context, req service.DeleteTodoRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	_, err := a.db.NamedExec(
		"DELETE FROM todo WHERE id = :id LIMIT 1",
		map[string]interface{}{
			"id": req.Id,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
