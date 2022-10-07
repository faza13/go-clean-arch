package uuid

import "github.com/google/uuid"

type IProvider interface {
	New() uuid.UUID
	GetString() string
}

type uuidProvider struct{}

func NewUuidProvider() IProvider {
	return &uuidProvider{}
}

func (u uuidProvider) New() uuid.UUID {
	res, _ := uuid.NewRandom()
	return res
}

func (u uuidProvider) GetString() string {
	res := u.New()
	return res.String()
}
