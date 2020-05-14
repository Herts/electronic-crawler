package routers

import (
	"github.com/Herts/electronic-crawler/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TaskController{})
	beego.Router("/addtask", &controllers.TaskController{}, "get:AddTaskPage")
	beego.Router("/api/task/add", &controllers.TaskController{}, "post:AddTask")
	beego.Router("/task", &controllers.TaskController{}, "get:ListPagesPage")
}
