package service

import (
	"context"

	"github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

type Pet struct {
	petAdapter PetAdapter
}

func NewPet(petAdapter PetAdapter) *Pet {
	return &Pet{petAdapter}
}

type CreatePetRequest struct {
	api.Pet
}

func (r CreatePetRequest) Validate() error {
	return nil
}

type GetPetsRequest struct {
	Offset int
	Limit  int
}

func (r GetPetsRequest) Validate() error {
	return nil
}

type GetPetRequest struct {
	Id int
}

func (r GetPetRequest) Validate() error {
	return nil
}

type UpdatePetRequest struct {
	Id   int
	Name *string
	Tag  *string
}

func (r UpdatePetRequest) Validate() error {
	return nil
}

type DeletePetRequest struct {
	Id int
}

func (r DeletePetRequest) Validate() error {
	return nil
}

func (s *Pet) Create(ctx context.Context, req CreatePetRequest) (*api.Pet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pet, err := s.petAdapter.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return pet, nil
}
func (s *Pet) List(ctx context.Context, req GetPetsRequest) ([]*api.Pet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pets, err := s.petAdapter.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return pets, nil
}
func (s *Pet) Get(ctx context.Context, req GetPetRequest) (*api.Pet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pet, err := s.petAdapter.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return pet, nil
}
func (s *Pet) Update(ctx context.Context, req UpdatePetRequest) (*api.Pet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	pet, err := s.petAdapter.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return pet, nil
}
func (s *Pet) Delete(ctx context.Context, req DeletePetRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	err := s.petAdapter.Delete(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
