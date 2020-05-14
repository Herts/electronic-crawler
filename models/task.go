package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	InitUrl      string `json:"initUrl"`
	DepthLimit   int    `json:"depthLimit,string"`
	Level        int    `json:"level,string"`
	Status       string
	KeyWords     string `json:"keyWords"`
	MaxProcesses int    `json:"maxProcesses,string"`
}

func GetAllTasks() (tasks []*Task) {
	db.Find(&tasks)
	return
}

func SaveTask(task *Task) {
	db.Save(task)
}
