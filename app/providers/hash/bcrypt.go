package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type IProvider interface {
	Make(needHash string) (string, error)
	Check(hashed string, compare string) bool
}

type provider struct {
}

func NewHashProvider() IProvider {
	return &provider{}
}

func (b provider) Make(needHash string) (string, error) {
	needHashByte := []byte(needHash)
	hashed, err := bcrypt.GenerateFromPassword(needHashByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (b provider) Check(hashed string, compare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(compare))
	if err != nil {
		return false
	}
	return true
}
