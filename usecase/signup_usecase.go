package usecase

import (
	"context"
	"time"

	"github.com/UxiT/rdp/domain"
	auth "github.com/UxiT/rdp/domain/auth"
	"github.com/UxiT/rdp/internal/tokenutil"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) auth.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetUserByLogin(c context.Context, login string) (domain.User, error) {
	return su.userRepository.GetByLogin(login)
}

func (su *signupUsecase) CreateAccessToken(c context.Context, user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(c context.Context, user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
