package token

import (
	"base/app/config"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type provider struct {
	jwtConfig *config.JWT
}

type IProvider interface {
	CreateToken(string, interface{}, time.Time) (string, error)
	CreateRefreshToken(string, time.Time) (string, error)
	ValidateToken(string) (interface{}, error)
	ValidateRefreshToken(string) (interface{}, error)
}

func NewTokenProvider(jwtConfig *config.JWT) IProvider {
	return &provider{
		jwtConfig: jwtConfig,
	}
}

func (j provider) CreateToken(tokenID string, data interface{}, exp time.Time) (string, error) {

	var mySigningKey = []byte(j.jwtConfig.JwtSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = tokenID
	claims["data"] = data
	claims["exp"] = exp.Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", errors.New("Tidak bisa generete token")
	}

	return tokenString, nil
}

func (j provider) CreateRefreshToken(tokenId string, exp time.Time) (string, error) {
	var mySigningKey = []byte(j.jwtConfig.RefreshSecret)
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = tokenId
	rtClaims["exp"] = exp.Unix()

	rt, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return rt, nil
}

func (j provider) ValidateToken(tokenString string) (interface{}, error) {
	var mySigningKey = []byte(j.jwtConfig.JwtSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["data"], nil

}

func (j provider) ValidateRefreshToken(tokenString string) (interface{}, error) {
	var mySigningKey = []byte(j.jwtConfig.RefreshSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["sub"], nil

}
