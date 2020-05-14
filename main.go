package main

import (
	"github.com/Herts/electronic-crawler/models"
	_ "github.com/Herts/electronic-crawler/routers"
	"github.com/astaxie/beego"
)

func init() {
	models.InitDB()
}

func main() {
	beego.Run()
}
