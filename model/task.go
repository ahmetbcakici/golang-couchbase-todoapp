package model

type Task struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}