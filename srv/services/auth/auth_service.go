package authService

import (
	"aramen-api/cmd/api/middleware"
	authModel "aramen-api/cmd/api/model/auth"
	"aramen-api/pkg/config"
	"aramen-api/srv/logs"
	authRepo "aramen-api/srv/repositories/auth"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthServiceInterface interface {
	CreateUser(user authModel.CreateUserModel) error
	RequestOtp(phoneNumber string) (*string, error)
	Login(login authModel.LoginModel) (*int, error)
}

type AuthService struct {
	cfg      *config.AppConfig
	authRepo authRepo.AuthRepo
}

func NewService(cfg *config.AppConfig, authRepo authRepo.AuthRepo) AuthService {
	return AuthService{cfg: cfg, authRepo: authRepo}
}

var (
	ErrGenerateTokenFailed = errors.New("Generate token failed")
)

func (s AuthService) generateToken(id string) (string, error) {
	// if err != nil {
	// 	logs.Error().Msgf("cannot generate token for %s: %s", id, err.Error())
	// 	return "", ErrGenerateTokenFailed
	// }

	claims := middleware.JWTClaim{ID: id}
	claims.ExpiresAt = time.Now().Add(time.Duration(s.cfg.JWT.ExpiredMinutes) * time.Minute).Unix()

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.cfg.JWT.SignedKey))
	if err != nil {
		logs.Error().Msgf("cannot generate token for %s: %s", id, err.Error())
		return "", ErrGenerateTokenFailed
	}

	return signedToken, nil
}

func (s AuthService) Login(login authModel.LoginModel) (*string, error) {
	userId, err := s.authRepo.Login(login)
	if err != nil {
		return nil, err
	}
	token, err := s.generateToken(fmt.Sprintf("%d", *userId))
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s AuthService) RequestOtp(phoneNumber string) (*string, error) {
	return s.authRepo.RequestOtp(phoneNumber)
}

func (s AuthService) CreateUser(user authModel.CreateUserModel) error {
	return s.authRepo.CreateUser(user)
}
