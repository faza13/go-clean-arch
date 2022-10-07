package services

import (
	"base/app/providers/hash"
	"base/enum"
	models2 "base/internal/auth/models"
	repositories2 "base/internal/auth/repositories"
	"base/internal/auth/resources"
	"base/internal/user/respositories"
	"base/utils/token"
	"base/utils/uuid"
	"errors"
	"time"
)

type IAuthService interface {
	Store(string, string, enum.Entity, string) (interface{}, error)
	Validate(string) (interface{}, error)
	GenerateNewToken(string) (interface{}, error)
}

type AuthService struct {
	userRepository         respositories.IUserRepository
	hashProvider           hash.IProvider
	tokenProvider          token.IProvider
	uuidProvider           uuid.IProvider
	refreshTokenRepository repositories2.IRefreshTokenRepository
}

func NewAuthService(userRepository respositories.IUserRepository, hashProvider hash.IProvider, token token.IProvider, uuidProvider uuid.IProvider, refreshTokenRepo repositories2.IRefreshTokenRepository) IAuthService {
	return &AuthService{
		userRepository:         userRepository,
		hashProvider:           hashProvider,
		tokenProvider:          token,
		uuidProvider:           uuidProvider,
		refreshTokenRepository: refreshTokenRepo,
	}
}

func (a AuthService) Store(username string, password string, entity enum.Entity, platform string) (interface{}, error) {
	user, err := a.userRepository.GetByUsername(username, entity)
	if err != nil {
		return nil, errors.New("Username tidak ditemukan")
	}

	validate := a.hashProvider.Check(user.Password, password)

	if !validate {
		return nil, errors.New("Username atau password anda salah")
	}

	data := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"entity":   1,
	}

	tokenId := a.uuidProvider.New()
	tokenIdStr := tokenId.String()

	token, err := a.tokenProvider.CreateToken(tokenIdStr, data, a.getExpAccessToken())

	if err != nil {
		return nil, err
	}

	expired := a.getExpRefreshToken()
	refreshToken, err := a.tokenProvider.CreateRefreshToken(tokenIdStr, expired)

	if err != nil {
		return nil, err
	}

	refreshTokenModel := models2.RefreshToken{
		ID:        uuid.ParseUUID(tokenIdStr),
		Uuid:      tokenIdStr,
		UserId:    user.ID,
		Platform:  platform,
		UserAgent: "",
		ExpiredAt: expired,
		Counter:   0,
	}

	err = a.refreshTokenRepository.Store(&refreshTokenModel)

	if err != nil {
		return nil, err
	}

	resp := resources.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpireIn:     expired.Unix(),
	}

	return resp, nil
}

func (a AuthService) Validate(token string) (interface{}, error) {
	res, err := a.tokenProvider.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a AuthService) GenerateNewToken(refreshToken string) (interface{}, error) {
	tokenId, err := a.tokenProvider.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	tokenIdStr := tokenId.(string)
	uuidStrTokenID := uuid.ParseUUID(tokenIdStr)
	refreshTokenModel, err := a.refreshTokenRepository.GetById(uuidStrTokenID)
	if err != nil {
		return nil, err
	}

	var today = time.Now()

	exp := today.After(refreshTokenModel.ExpiredAt)

	if exp {
		return nil, errors.New("refresh token telah kadaluarsa")
	}

	counter := refreshTokenModel.Counter + 1
	refreshTokenMap := make(map[string]interface{})
	refreshTokenMap["counter"] = counter

	if counter > 5 {
		newTokenId := a.uuidProvider.New()
		refreshTokenMap["counter"] = 0
		refreshTokenMap["id"] = uuid.BinaryUUID(newTokenId)
		tokenIdStr = newTokenId.String()
		refreshTokenMap["uuid"] = tokenIdStr
	}

	_, err = a.refreshTokenRepository.Update(uuidStrTokenID, refreshTokenMap)
	if err != nil {
		return nil, err
	}

	if refreshTokenMap["uuid"] != "" {
		expired := a.getExpRefreshToken()
		refreshToken, err = a.tokenProvider.CreateRefreshToken(tokenIdStr, expired)
	}
	user, err := a.userRepository.GetById(refreshTokenModel.UserId)

	data := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
	}
	expireIn := a.getExpAccessToken()
	token, err := a.tokenProvider.CreateToken(tokenIdStr, data, expireIn)

	resp := resources.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpireIn:     expireIn.Unix(),
	}

	return resp, nil
}

func (a AuthService) getExpAccessToken() time.Time {
	return time.Now().Add(time.Hour * 15)
}

func (a AuthService) getExpRefreshToken() time.Time {
	return time.Now().AddDate(1, 0, 0)
}
