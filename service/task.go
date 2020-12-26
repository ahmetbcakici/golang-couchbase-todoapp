package service

import (
	"encoding/json"
	"github.com/couchbase/gocb"
	"github.com/labstack/echo"
	"go-todoapp/model"
	"io/ioutil"
)

type ITaskService interface {
	GetTasks() ([]model.Task, error)
	GetTaskById(ctx echo.Context) (*model.Task, error)
	SaveNewTask(ctx echo.Context) (*model.Task, error)
	UpdateTaskById(ctx echo.Context) (*model.Task, error)
	RemoveTaskById(ctx echo.Context) (*model.Task, error)
}

type TaskService struct {
	Cluster gocb.Cluster
}

func (s TaskService) GetTasks() ([]model.Task, error) {
	var tasks []model.Task
	var task model.Task

	bucket, _ := s.Cluster.OpenBucket("task", "")

	query := gocb.NewN1qlQuery("SELECT id,name,status FROM task")
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		return nil, err
	}

	for rows.Next(&task) {
		tasks = append(tasks, task)
		task = model.Task{}
	}

	return tasks, nil
}

func (s TaskService) GetTaskById(ctx echo.Context) (*model.Task, error) {
	taskId := ctx.Param("taskId")
	task := model.Task{}

	bucket, _ := s.Cluster.OpenBucket("task", "")

	_, err := bucket.Get(taskId, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s TaskService) SaveNewTask(ctx echo.Context) (*model.Task, error) {
	task := model.Task{
		Id: "2",
	}

	bucket, _ := s.Cluster.OpenBucket("task", "")
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		return nil, err
	}

	_, err = bucket.Insert(task.Id,
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

func (s TaskService) UpdateTaskById(ctx echo.Context) (*model.Task, error) {
	task := model.Task{}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		return nil, err
	}

	bucket, _ := s.Cluster.OpenBucket("task", "")
	_, err = bucket.Replace(task.Id,
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

func (s TaskService) RemoveTaskById(ctx echo.Context) (*model.Task, error) {
	taskId := ctx.Param("taskId")

	bucket, _ := s.Cluster.OpenBucket("task", "")
	_, err := bucket.Remove(taskId, 0)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
