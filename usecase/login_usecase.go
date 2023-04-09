package usecase

import (
	"context"
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

func (lu *loginUsecase) GetUserByLogin(c context.Context, login string) (domain.User, error) {
	return lu.userRepository.GetByLogin(login)
}

func (lu *loginUsecase) CreateAccessToken(c context.Context, user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(c context.Context, user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
