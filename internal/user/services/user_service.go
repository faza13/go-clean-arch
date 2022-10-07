package services

import (
	"base/app/providers/hash"
	"base/internal/user/models"
	"base/internal/user/resources"
	"base/internal/user/respositories"
)

type IUserService interface {
	GetAll() []*models.User
	Update(id uint64, user *models.User) *models.User
	Store(user *resources.UserStoreRequest) (*models.User, error)
	ChangePassword(id uint64, newPassword string) error
}

type UserService struct {
	userRepository respositories.IUserRepository
	hashProvider   hash.IProvider
}

func NewUserService(userRepository respositories.IUserRepository, hashProvider hash.IProvider) IUserService {
	return &UserService{
		userRepository: userRepository,
		hashProvider:   hashProvider,
	}
}

func (s UserService) GetAll() []*models.User {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Update(id uint64, user *models.User) *models.User {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Store(user *resources.UserStoreRequest) (*models.User, error) {

	passwordHashed, err := s.hashProvider.Make(user.Password)

	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Password: passwordHashed,
		Entity:   0,
	}

	err = s.userRepository.Store(&newUser)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (s UserService) ChangePassword(id uint64, newPassword string) error {
	passwordHash, err := s.hashProvider.Make(newPassword)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"password": passwordHash,
	}
	err = s.userRepository.Update(id, data)

	if err != nil {
		return err
	}

	return nil
}
