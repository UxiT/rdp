package route

import (
	"time"

	"github.com/UxiT/rdp/api/controller"
	"github.com/UxiT/rdp/bootstrap"
	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/repository"
	"github.com/UxiT/rdp/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewUserTaskRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	utr := repository.NewUserTaskRepository(db)
	tr := repository.NewTaskRepository(db)
	utc := &controller.TaskController{
		UserTaskModel: usecase.NewUserTaskUsecase(utr, timeout),
		TaskModel:     usecase.NewTaskUsecase(tr, timeout),
		Env:           env,
	}

	group.GET("/task/by-course", utc.GetByCourse)
	group.GET("/task/:id", utc.GetTask)

	group.OPTIONS("/task/by-course", cors.Default())
	group.OPTIONS("/task/:id", cors.Default())
}
