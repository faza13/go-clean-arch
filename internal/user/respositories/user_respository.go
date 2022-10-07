package respositories

import (
	"base/app/providers/database"
	"base/enum"
	"base/internal/user/models"
)

type UserRepository struct {
	DBProvider *database.Provider
}

type IUserRepository interface {
	Store(user *models.User) error
	GetByUsername(username string, entity enum.Entity) (*models.User, error)
	GetById(id uint64) (*models.User, error)
	Update(id uint64, data map[string]interface{}) error
}

func NewUserRepository(dbProvider *database.Provider) IUserRepository {
	return &UserRepository{
		DBProvider: dbProvider,
	}
}

func (u UserRepository) Store(user *models.User) error {
	result := u.DBProvider.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserRepository) GetByUsername(username string, entity enum.Entity) (*models.User, error) {
	var user models.User
	result := u.DBProvider.DB.First(&user, "username = ? AND entity = ?", username, entity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u UserRepository) GetById(id uint64) (*models.User, error) {
	var user models.User
	result := u.DBProvider.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u UserRepository) Update(id uint64, data map[string]interface{}) error {
	if err := u.DBProvider.DB.Model(models.User{}).Where("id=?", id).UpdateColumns(data).Error; err != nil {
		return err
	}
	return nil
}
