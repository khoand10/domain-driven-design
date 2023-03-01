package service

import (
	"context"
	"domain-driven-design/config"
	"domain-driven-design/domain/repository"
	"domain-driven-design/pkg/jwt"
	"domain-driven-design/pkg/utils"
	"errors"
)

type (
	LoginReq struct {
		Email    string
		Password string
	}
)

type AuthService interface {
	Login(ctx context.Context, loginReq *LoginReq) (*jwt.TokenInfo, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (*jwt.TokenInfo, error)
}

type authService struct {
	datastore repository.UserRepository
	config    *config.Config
}

func NewAuthService(datastore repository.UserRepository, config *config.Config) AuthService {
	return &authService{
		datastore: datastore,
		config:    config,
	}
}

func (a authService) RefreshAccessToken(ctx context.Context, refreshToken string) (*jwt.TokenInfo, error) {
	claims, err := jwt.VerifyToken(refreshToken, a.config.JwtSecretKey)
	if err != nil {
		return nil, err
	}

	userFound, err := a.datastore.GetByID(ctx, claims.UserId)
	if err != nil {
		return nil, errors.New("user is not exist")
	}
	newToken, err := jwt.CreateJWT(userFound, a.config.JwtSecretKey, a.config.TokenExpirationHour)
	if err != nil {
		return nil, err
	}

	tokenInfo := &jwt.TokenInfo{
		Token:        newToken,
		RefreshToken: refreshToken,
	}

	return tokenInfo, nil
}

func (a authService) Login(ctx context.Context, loginReq *LoginReq) (*jwt.TokenInfo, error) {
	userFound, err := a.datastore.GetByEmail(ctx, loginReq.Email)
	if err != nil {
		return nil, errors.New("user is not exist")
	}

	if err := utils.ComparePassword(userFound.Password, loginReq.Password); err != nil {
		return nil, errors.New("email or password is incorrect")
	}

	token, err := jwt.CreateJWT(userFound, a.config.JwtSecretKey, a.config.TokenExpirationHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateJWT(userFound, a.config.JwtSecretKey, a.config.RefreshTokenExpirationHour)
	if err != nil {
		return nil, err
	}
	tokenInfo := &jwt.TokenInfo{
		Token:        token,
		RefreshToken: refreshToken,
	}

	return tokenInfo, nil
}
