package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/UxiT/rdp/domain"
	auth "github.com/UxiT/rdp/domain/auth"
	"github.com/UxiT/rdp/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) auth.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByLogin(c context.Context, login string) (*domain.User, error) {
	users, err := lu.userRepository.GetByField("login", login)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if len(users) == 0 {
		return nil, errors.New("invalid credentials")
	}

	return &users[0], nil
}

func (lu *loginUsecase) CreateAccessToken(c context.Context, user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(c context.Context, user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
