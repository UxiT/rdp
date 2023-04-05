package route

import (
	"time"

	"github.com/UxiT/rdp/api/controller"
	"github.com/UxiT/rdp/bootstrap"
	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/domain"
	"github.com/UxiT/rdp/repository"
	"github.com/UxiT/rdp/usecase"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
