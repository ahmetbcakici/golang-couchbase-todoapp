package controller

import (
	"github.com/couchbase/gocb"
	"github.com/labstack/echo"
	"go-todoapp/service"
	"net/http"
)

type ITaskController interface {
	GetTasks(ctx echo.Context) error
}

type TaskController struct {
	Cluster *gocb.Cluster
	taskService service.TaskService
}

func (c *TaskController) GetTasks(ctx echo.Context) error {
	tasks, _ := service.TaskService{Cluster: *c.Cluster}.GetTasks()
	return ctx.JSON(http.StatusOK, tasks)
}


func (c *TaskController) GetTaskById(ctx echo.Context) error {
	task, _ := service.TaskService{Cluster: *c.Cluster}.GetTaskById(ctx)
	return ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) SaveNewTask(ctx echo.Context) error {
	service.TaskService{Cluster: *c.Cluster}.SaveNewTask(ctx)
	return ctx.String(200, "ok")
}

func (c *TaskController) UpdateTaskById(ctx echo.Context) error {
	service.TaskService{Cluster: *c.Cluster}.UpdateTaskById(ctx)
	return ctx.String(200, "ok")
}

func (c *TaskController) RemoveTaskById(ctx echo.Context) error {
	service.TaskService{Cluster: *c.Cluster}.RemoveTaskById(ctx)
	return ctx.String(200, "ok")
}
