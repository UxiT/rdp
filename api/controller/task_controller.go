package controller

import (
	"fmt"
	"net/http"

	"github.com/UxiT/rdp/bootstrap"
	"github.com/UxiT/rdp/domain"
	task "github.com/UxiT/rdp/domain/tasks"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	UserTaskModel task.UserTaskModel
	TaskModel     task.TaskModel
	Env           *bootstrap.Env
}

func (tc *TaskController) GetByCourse(c *gin.Context) {
	userId := c.GetString("x-user-id")
	profile := c.GetString("x-user-profile")
	courseId, ok := c.GetQuery("course_id")

	if !ok {
		fmt.Errorf("Invalid courseId: %v", courseId)
	}

	var tasksR []task.UserTask

	if profile == "student" {
		tasks, err := tc.UserTaskModel.GetByCourse(c, userId, courseId)
		tasksR = tasks

		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, tasksR)
	} else {
		tasks, err := tc.UserTaskModel.GetByCourse(c, userId, courseId)
		tasksR = tasks

		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, tasksR)
	}
}

func (tc *TaskController) GetTask(c *gin.Context) {
	userId := c.GetString("x-user-id")
	profile := c.GetString("x-user-profile")
	taskId := c.Param("id")

	if profile == "student" {
		task, err := tc.TaskModel.Read(c, userId, taskId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, task)
	} else {
		task, err := tc.TaskModel.Read(c, userId, taskId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}
