package controller

import (
	"github.com/labstack/echo"
	"go-todoapp/service"
	"net/http"
)

func GetTasks(ctx echo.Context) error {
	tasks, _ := service.GetTasks()
	return ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById(ctx echo.Context) error {
	task, _ := service.GetTaskById(ctx)
	return ctx.JSON(http.StatusOK, task)
}

func SaveNewTask(ctx echo.Context) error {
	service.SaveNewTask(ctx)
	return ctx.String(200, "ok")
}

func UpdateTaskById(ctx echo.Context) error {
	service.UpdateTaskById(ctx)
	return ctx.String(200, "ok")
}

func RemoveTaskById(ctx echo.Context) error {
	service.RemoveTaskById(ctx)
	return ctx.String(200, "ok")
}
