//go:build db
// +build db

package adapter

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

func TestDB_Todo_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req service.CreateTodoRequest
	}
	tests := []struct {
		name    string
		a       *Todo
		args    args
		want    *api.Todo
		wantErr bool
	}{
		{
			name: "登録できること",
			args: args{
				ctx: context.TODO(),
				req: service.CreateTodoRequest{
					Todo: api.Todo{
						Subject: "subject a",
						Body:    "body",
					},
				},
			},
			want: &api.Todo{
				Subject: "subject a",
				Body:    "body",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewTodo(db, id)
			got, err := a.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := []cmp.Option{
				cmpopts.IgnoreFields(api.Todo{}, "Id", "CreatedAt", "UpdatedAt"),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Todo.Create(got , want) \n%s", diff)
			}
		})
	}
}

func TestTodo_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.GetTodoRequest
	}
	tests := []struct {
		name string

		args        args
		want        *api.Todo
		wantErr     bool
		fixtureFunc func()
	}{
		{
			name: "登録できること",
			args: args{
				ctx: context.TODO(),
				req: service.GetTodoRequest{
					Id: "cmfj5jhk5epvhvf7gqn0",
				},
			},
			want: &api.Todo{
				Id:      "cmfj5jhk5epvhvf7gqn0",
				Subject: "subject a",
				Body:    "body",
			},
			wantErr: false,
			fixtureFunc: func() {
				db.MustExec(`INSERT INTO todo (id, subject, body) VALUES ("cmfj5jhk5epvhvf7gqn0","subject a","body")`)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fixtureFunc()

			a := NewTodo(db, id)
			got, err := a.Get(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := []cmp.Option{
				cmpopts.IgnoreFields(api.Todo{}, "Id", "CreatedAt", "UpdatedAt"),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Todo.Create(got , want) \n%s", diff)
			}
		})
	}
}
