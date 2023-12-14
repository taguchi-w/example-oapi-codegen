package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
)

func TestPet_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req CreatePetRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *api.Pet
		wantErr bool
		mocks   map[string]interface{}
	}{
		{
			name: "登録できること",
			args: args{
				ctx: context.TODO(),
				req: CreatePetRequest{
					Pet: api.Pet{
						Name: "test",
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
				"petAdapter.Create.pet": &api.Pet{
					Id:   "1",
					Name: "test",
					Tag:  util.P("tag"),
				},
				"petAdapter.Create.err": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPetAdapter := NewMockPetAdapter(ctrl)
			mockPetAdapter.EXPECT().Create(tt.args.ctx, tt.args.req).Return(
				tt.mocks["petAdapter.Create.pet"],
				tt.mocks["petAdapter.Create.err"],
			)
			s := &Pet{
				petAdapter: mockPetAdapter,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
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
