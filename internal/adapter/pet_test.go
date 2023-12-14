package adapter

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/taguchi-w/example-oapi-codegen/internal/service"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
)

func TestPet_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req service.CreatePetRequest
	}
	tests := []struct {
		name    string
		a       *Pet
		args    args
		want    *api.Pet
		wantErr bool
		mocks   map[string]interface{}
	}{
		{
			name: "登録できること",
			args: args{
				ctx: context.TODO(),
				req: service.CreatePetRequest{
					Pet: api.Pet{
						Name: "test",
						Tag:  util.P("tag"),
					},
				},
			},
			want: &api.Pet{
				Id:   "1",
				Name: "test",
				Tag:  util.P("tag"),
			},
			wantErr: false,
			mocks: map[string]interface{}{
				"db.NamedExec.result": nil,
				"db.NamedExec.err":    nil,
				"id.Generate.id":      "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewMockDBAdapter(ctrl)
			id := NewMockIDGenerator(ctrl)
			db.EXPECT().NamedExec(gomock.Any(), gomock.Any()).Return(
				tt.mocks["db.NamedExec.result"],
				tt.mocks["db.NamedExec.err"],
			)
			id.EXPECT().Generate().Return(tt.mocks["id.Generate.id"])

			a := &Pet{db, id}
			got, err := a.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pet.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Pet.Create(got , want) \n%s", diff)
			}
		})
	}
}
