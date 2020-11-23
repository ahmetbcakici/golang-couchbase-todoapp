package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go-todoapp/controller"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {
	e := echo.New()

	e.GET("/", controller.GetTasks)
	e.GET("/:taskId", controller.GetTaskById)
	e.POST("/", controller.SaveNewTask)
	e.PATCH("/", controller.UpdateTaskById)
	e.DELETE("/:taskId", controller.RemoveTaskById)

	e.Start(":8080")
}
