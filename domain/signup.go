package domain

type SignupRequest struct {
	Name      string `form:"name" binding:"required"`
	Last_Name string `form:"last_name"`
	Login     string `form:"login" binding:"required"`
	Password  string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(user *User) error
	GetUserByLogin(email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
