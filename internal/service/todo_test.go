package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

func TestTodo_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req CreateTodoRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *api.Todo
		wantErr bool
		mocks   map[string]interface{}
	}{
		{
			name: "登録できること",
			args: args{
				ctx: context.TODO(),
				req: CreateTodoRequest{
					Todo: api.Todo{
						Subject: "subject a",
						Body:    "body",
					},
				},
			},
			want: &api.Todo{
				Id:      "1",
				Subject: "subject a",
				Body:    "body",
			},
			wantErr: false,
			mocks: map[string]interface{}{
				"todoAdapter.Create.pet": &api.Todo{
					Id:      "1",
					Subject: "subject a",
					Body:    "body",
				},
				"todoAdapter.Create.err": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTodoAdapter := NewMockTodoAdapter(ctrl)
			mockTodoAdapter.EXPECT().Create(tt.args.ctx, tt.args.req).Return(
				tt.mocks["todoAdapter.Create.pet"],
				tt.mocks["todoAdapter.Create.err"],
			)
			s := &Todo{
				todoAdapter: mockTodoAdapter,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Todo.Create(got , want) \n%s", diff)
			}

		})
	}
}
