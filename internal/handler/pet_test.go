package handler

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewPet(t *testing.T) {
	type args struct {
		pet PetService
	}
	tests := []struct {
		name string
		args args
		want *Pet
	}{
		{
			name: "",
			args: args{},
			want: &Pet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPet(tt.args.pet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPet_GetPets(t *testing.T) {
	type fields struct {
		Pet PetService
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Pet{
				Pet: tt.fields.Pet,
			}
			if err := h.GetPets(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Pet.GetPets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPet_PostPets(t *testing.T) {
	type fields struct {
		Pet PetService
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Pet{
				Pet: tt.fields.Pet,
			}
			if err := h.PostPets(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Pet.PostPets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPet_UpdatePetPartial(t *testing.T) {
	type fields struct {
		Pet PetService
	}
	type args struct {
		ctx   echo.Context
		petId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Pet{
				Pet: tt.fields.Pet,
			}
			if err := h.UpdatePetPartial(tt.args.ctx, tt.args.petId); (err != nil) != tt.wantErr {
				t.Errorf("Pet.UpdatePetPartial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPet_DeletePet(t *testing.T) {
	type fields struct {
		Pet PetService
	}
	type args struct {
		ctx   echo.Context
		petId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Pet{
				Pet: tt.fields.Pet,
			}
			if err := h.DeletePet(tt.args.ctx, tt.args.petId); (err != nil) != tt.wantErr {
				t.Errorf("Pet.DeletePet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
