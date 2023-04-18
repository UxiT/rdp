package route

import (
	"time"

	"github.com/UxiT/rdp/api/controller"
	"github.com/UxiT/rdp/bootstrap"
	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/domain"
	"github.com/UxiT/rdp/repository"
	"github.com/UxiT/rdp/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCoursesRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	cr := repository.NewCourseRepository(db, domain.CollectionUser)
	cc := &controller.CoursesController{
		CoursesUsecase: usecase.NewCoursesUsecase(cr, timeout),
		Env:            env,
	}
	group.GET("/courses", cc.FetchByUser)
	group.POST("/courses/create", cc.Create)
	group.POST("/courses/task-attach", cc.AttachTask)

	group.OPTIONS("courses/create", cors.Default())
	group.OPTIONS("/courses", cors.Default())
}
