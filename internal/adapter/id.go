package adapter

import "github.com/rs/xid"

type XIDGenerator struct{}

func (a *XIDGenerator) Generate() string {
	return xid.New().String()
}

func NewXIDGenerator() *XIDGenerator {
	return &XIDGenerator{}
}
