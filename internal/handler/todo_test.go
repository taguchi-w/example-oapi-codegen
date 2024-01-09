package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
)

func TestTodo_PostTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		wantErr    bool
		want       *api.Todo
		wantStatus int
		mocks      map[string]interface{}
	}{
		{
			name:    "",
			wantErr: false,
			mocks: map[string]interface{}{
				"pet.Create.result": &api.Todo{
					Id:      "1",
					Subject: "subject a",
					Body:    util.P("body"),
				},
				"pet.Create.err": nil,
			},
			want: &api.Todo{
				Id:      "1",
				Subject: "subject a",
				Body:    util.P("body"),
			},
			wantStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			pet := NewMockTodoService(ctrl)
			pet.EXPECT().Create(gomock.Any(), gomock.Any()).Return(
				tt.mocks["pet.Create.result"].(*api.Todo),
				tt.mocks["pet.Create.err"],
			)
			h := &Todo{Todo: pet}
			if err := h.PostTodos(ctx); (err != nil) != tt.wantErr {
				t.Errorf("Todo.PostTodos() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, rec.Result().StatusCode, tt.wantStatus)

			if !tt.wantErr {
				var got api.Todo
				if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if diff := cmp.Diff(&got, tt.want); diff != "" {
					t.Errorf("Todo.PostTodos(got , want) \n%s", diff)
				}
			}
		})
	}
}
