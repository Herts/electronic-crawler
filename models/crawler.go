package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/axgle/mahonia"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/html/charset"
)

type Page struct {
	gorm.Model
	Url       string `gorm:"primary_key;varchar(500)"`
	Title     string
	Level     int
	ParentUrl string
	TaskID    uint
}

type UrlsKnown struct {
	gorm.Model
	Url string `gorm:"primary_key;varchar(500)"`
}

func StartTask(task *Task) {

	collector := &Collector{}
	collector.init()

	req := gorequest.New()
	targetUrl := task.InitUrl
	collector.Urls[targetUrl] = task.Level
	_, body, errs := req.Get(targetUrl).End()

	collector.Parse(targetUrl, body, errs, task.Level, "", task.ID)

	newUrls := collector.GetAllUnvisited()
	logs.Debug(len(newUrls))
	for _, u := range newUrls {
		_, body, errs := req.Get(u).End()
		collector.Parse(u, body, errs, task.Level, targetUrl, task.ID)
	}
	task.Status = "Ended"
	SaveTask(task)
}

func decoderConvert(name string, body string) string {
	return mahonia.NewDecoder(name).ConvertString(body)
}

func determineEncoding(body string) string {
	byteBody := []byte(body)
	if len(byteBody) > 1024 {
		byteBody = byteBody[:1024]
	}

	_, name, _ := charset.DetermineEncoding(byteBody, "")
	return name
}

func GetPagesByTaskID(taskID uint) (pages []*Page) {
	db.Where(&Page{TaskID: taskID}).Order("level").Find(&pages)
	return
}
