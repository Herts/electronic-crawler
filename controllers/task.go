package controllers

import (
	"encoding/json"
	"github.com/Herts/electronic-crawler/models"
	"github.com/astaxie/beego"
	"net/url"
)

type TaskController struct {
	beego.Controller
}

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *TaskController) Get() {
	tasks := models.GetAllTasks()
	c.Data["Title"] = "Tasks"
	c.Layout = "layout.html"
	c.TplName = "tasks.html"
	c.Data["tasks"] = tasks
}

func (c *TaskController) AddTaskPage() {
	c.Data["Title"] = "Add a task"
	c.Layout = "layout.html"
	c.TplName = "edittask.html"
}

func (c *TaskController) AddTask() {
	var task models.Task
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &task)
	if err != nil {
		c.Data["json"] = response{Message: "json.Unmarshal " + err.Error()}
		c.ServeJSON()
		return
	}

	_, err = url.Parse(task.InitUrl)
	if err != nil {
		c.Data["json"] = response{Message: "json.Unmarshal" + err.Error()}
		c.ServeJSON()
		return
	}

	task.Status = "Running"
	models.SaveTask(&task)
	go models.StartTask(&task)

	c.Data["json"] = response{Message: "Task was added successfully"}
	c.ServeJSON()
}

func (c *TaskController) ListPagesPage() {
	c.Layout = "layout.html"
	c.TplName = "pages.html"
	id, err := c.GetUint64("id")
	if err != nil {
		c.Data["json"] = response{Message: "json.Unmarshal" + err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["pages"] = models.GetPagesByTaskID(uint(id))
}
