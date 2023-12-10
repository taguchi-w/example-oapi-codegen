package adapter

import "github.com/taguchi-w/example-oapi-codegen/internal/service"

func New() service.Adapters {
	return service.Adapters{
		PetManager: NewPet(),
	}
}
