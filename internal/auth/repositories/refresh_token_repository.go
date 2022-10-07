package repositories

import (
	"base/app/providers/database"
	"base/internal/auth/models"
	"base/utils/uuid"
	errors "errors"
	"gorm.io/gorm"
)

type refreshTokenRespository struct {
	dbProvider *database.Provider
}

type IRefreshTokenRepository interface {
	Store(data *models.RefreshToken) error
	GetById(id uuid.BinaryUUID) (*models.RefreshToken, error)
	Update(id uuid.BinaryUUID, toBeUpdate map[string]interface{}) (*models.RefreshToken, error)
}

func NewRefreshTokenRepositories(dbProvider *database.Provider) IRefreshTokenRepository {
	return &refreshTokenRespository{
		dbProvider: dbProvider,
	}
}

func (r refreshTokenRespository) Store(data *models.RefreshToken) error {

	res := r.dbProvider.DB.Create(&data)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r refreshTokenRespository) GetById(id uuid.BinaryUUID) (*models.RefreshToken, error) {
	refreshToken := models.RefreshToken{}

	res := r.dbProvider.DB.Where("id = ?", id).First(&refreshToken)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	return &refreshToken, nil
}

func (r refreshTokenRespository) Update(id uuid.BinaryUUID, toBeUpdate map[string]interface{}) (*models.RefreshToken, error) {
	newRefreshToken := models.RefreshToken{}
	err := r.dbProvider.DB.Model(&newRefreshToken).Where("id=?", id).Updates(toBeUpdate).Error
	if err != nil {
		return nil, err
	}
	return &newRefreshToken, nil
}
