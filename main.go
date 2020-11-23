package main

import (
	"github.com/labstack/echo"
	"go-todoapp/controller"
)

func main() {
	e := echo.New()

	e.GET("/", controller.GetTasks)
	e.GET("/:taskId", controller.GetTaskById)
	e.POST("/", controller.SaveNewTask)
	e.PATCH("/", controller.UpdateTaskById)
	e.DELETE("/:taskId", controller.RemoveTaskById)

	e.Start(":8080")
}
