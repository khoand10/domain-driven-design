package service

import (
	"context"
	"domain-driven-design/config"
	"domain-driven-design/domain/repository"
	"domain-driven-design/pkg/jwt"
	"domain-driven-design/pkg/utils"
	"errors"
	"time"
)

type (
	LoginReq struct {
		Email    string
		Password string
	}
	LoginRes struct {
		Token                  string    `json:"token"`
		RefreshToken           string    `json:"refresh_token"`
		RefreshTokenExpireTime time.Time `json:"refresh_token_expire_time"`
		TokenExpireTime        time.Time `json:"token_expire_time"`
	}

	RefreshReq struct {
		RefreshToken string `json:"refresh_token"`
	}

	RefreshRes struct {
		Token                  string    `json:"token"`
		RefreshToken           string    `json:"refresh_token"`
		RefreshTokenExpireTime time.Time `json:"refresh_token_expire_time"`
		TokenExpireTime        time.Time `json:"token_expire_time"`
	}
)

type AuthService interface {
	Login(ctx context.Context, loginReq *LoginReq) (*LoginRes, error)
	RefreshAccessToken(ctx context.Context, refreshReq *RefreshReq) (*RefreshRes, error)
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

func (a authService) RefreshAccessToken(ctx context.Context, refreshReq *RefreshReq) (*RefreshRes, error) {
	claims, err := jwt.VerifyToken(refreshReq.RefreshToken, a.config.JwtSecretKey)
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

	res := &RefreshRes{
		Token:                  newToken,
		RefreshToken:           refreshReq.RefreshToken,
		TokenExpireTime:        time.Now().Add(time.Hour * time.Duration(a.config.TokenExpirationHour)),
		RefreshTokenExpireTime: time.Now().Add(time.Hour * time.Duration(a.config.RefreshTokenExpirationHour)),
	}

	return res, nil
}

func (a authService) Login(ctx context.Context, loginReq *LoginReq) (*LoginRes, error) {
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
	loginRes := &LoginRes{
		Token:                  token,
		RefreshToken:           refreshToken,
		TokenExpireTime:        time.Now().Add(time.Hour * time.Duration(a.config.TokenExpirationHour)),
		RefreshTokenExpireTime: time.Now().Add(time.Hour * time.Duration(a.config.RefreshTokenExpirationHour)),
	}

	return loginRes, nil
}
