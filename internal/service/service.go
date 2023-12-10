package service

type Services struct {
	Pet *Pet
}

type Adapters struct {
	PetManager PetManager
}

func New(adapters Adapters) Services {
	return Services{
		Pet: NewPet(adapters.PetManager),
	}
}
