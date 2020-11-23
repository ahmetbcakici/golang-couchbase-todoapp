package service

import (
	"encoding/json"
	"github.com/couchbase/gocb"
	"github.com/labstack/echo"
	"go-todoapp/common"
	"go-todoapp/model"
	"io/ioutil"
)

func GetTasks() ([]model.Task, error) {
	var tasks []model.Task
	var task model.Task

	query := gocb.NewN1qlQuery("SELECT id,name,status FROM task")
	rows, err := common.Bucket().ExecuteN1qlQuery(query, nil)
	if err != nil {
		return nil, err
	}

	for rows.Next(&task) {
		tasks = append(tasks, task)
		task = model.Task{}
	}

	return tasks, nil
}

func GetTaskById(ctx echo.Context) (*model.Task, error) {
	taskId := ctx.Param("taskId")
	task := model.Task{}

	_, err := common.Bucket().Get(taskId, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func SaveNewTask(ctx echo.Context) (*model.Task, error) {
	task := model.Task{
		Id: "2",
	}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		return nil, err
	}

	_, err = common.Bucket().Insert(task.Id,
		model.Task{
			Id:     task.Id,
			Name:   task.Name,
			Status: task.Status,
		}, 0)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTaskById(ctx echo.Context) (*model.Task, error) {
	task := model.Task{}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		return nil, err
	}

	_, err = common.Bucket().Replace(task.Id,
		model.Task{
			Id:     task.Id,
			Name:   task.Name,
			Status: task.Status,
		}, 0, 0)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func RemoveTaskById(ctx echo.Context) (*model.Task, error) {
	taskId := ctx.Param("taskId")

	_, err := common.Bucket().Remove(taskId, 0)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
