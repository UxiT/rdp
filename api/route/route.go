package route

import (
	"time"

	"github.com/UxiT/rdp/api/middleware"
	"github.com/UxiT/rdp/bootstrap"
	"github.com/UxiT/rdp/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db db.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	publicRouter.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
	}))
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
	}))
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewCoursesRouter(env, timeout, db, protectedRouter)
	NewUserTaskRouter(env, timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)
}
