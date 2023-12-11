//go:generate mockgen -source=adapter.go -destination=mocks/mock_adapter.go -package=mock
package adapter

type Adapters struct {
	Pet *Pet
}

func New(db DBAdapter) Adapters {
	return Adapters{
		Pet: NewPet(),
	}
}

type DBAdapter interface {
}
