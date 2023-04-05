package route

// import (
// 	"time"

// 	"github.com/UxiT/rdp/api/controller"
// 	"github.com/UxiT/rdp/bootstrap"
// 	"github.com/UxiT/rdp/db"
// 	"github.com/UxiT/rdp/domain"
// 	"github.com/UxiT/rdp/repository"
// 	"github.com/UxiT/rdp/usecase"
// 	"github.com/gin-gonic/gin"
// )

// func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
// 	tr := repository.NewTaskRepository(db, domain.CollectionTask)
// 	tc := &controller.TaskController{
// 		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
// 	}
// 	// group.GET("/task", tc.Fetch)
// 	group.POST("/task", tc.Create)
// }
