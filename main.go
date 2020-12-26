package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go-todoapp/common"
	"go-todoapp/controller"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {
	e := echo.New()

	cluster := common.Cluster()
	taskController := controller.TaskController{Cluster: cluster}

	e.GET("/a", func(context echo.Context) error {
		return context.String(200, "guzel")
	})

	e.GET("/", taskController.GetTasks)
	e.GET("/:taskId", taskController.GetTaskById)
	e.POST("/", taskController.SaveNewTask)
	e.PATCH("/", taskController.UpdateTaskById)
	e.DELETE("/:taskId", taskController.RemoveTaskById)

	e.Start(":8080")
}