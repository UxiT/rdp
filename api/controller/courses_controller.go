package controller

import (
	"fmt"
	"net/http"

	"github.com/UxiT/rdp/bootstrap"
	"github.com/UxiT/rdp/domain"
	courses "github.com/UxiT/rdp/domain/course"
	"github.com/gin-gonic/gin"
)

type CoursesController struct {
	CoursesUsecase courses.CourseModel
	Env            *bootstrap.Env
}

func (cc *CoursesController) FetchByUser(c *gin.Context) {
	userID := c.GetString("x-user-id")

	fmt.Printf("\nuser_ID: %s\n", userID)
	courses, err := cc.CoursesUsecase.FetchByUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (cc *CoursesController) Create(c *gin.Context) {
	// userID := c.GetString("x-user-id")

	var request courses.CreateCourseRequest
	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.CoursesUsecase.Create(c, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, "")
}

func (cc *CoursesController) AttachTask(c *gin.Context) {
	// userId := c.GetString("x-user-id")

	var request courses.AttachTaskRequest
	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

}
