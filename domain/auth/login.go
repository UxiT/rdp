package auth

import (
	"context"

	"github.com/UxiT/rdp/domain"
)

type LoginRequest struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByLogin(c context.Context, email string) (*domain.User, error)
	CreateAccessToken(c context.Context, user *domain.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(c context.Context, user *domain.User, secret string, expiry int) (refreshToken string, err error)
}
