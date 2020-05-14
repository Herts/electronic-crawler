package models

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type Collector struct {
	Urls        map[string]int
	VisitedUrls map[string]struct{}
	Keywords    []string
}

func (c *Collector) init() {
	c.Urls = make(map[string]int)
	c.VisitedUrls = make(map[string]struct{})
	c.Keywords = []string{"输变电工程", "高压架空线路", "高压输电线路", "高压线", "杆塔", "铁塔", "变电站", "换流站", "辐射"}
}

func (c *Collector) GetAllUnvisited() []string {
	var urls []string
	for u := range c.Urls {
		urls = append(urls, u)
	}
	return urls
}

func (c *Collector) Parse(targetUrl string, body string, errs []error, parentLevel int, parentUrl string, taskID uint) {
	if errs != nil {
		log.Println("accessing", targetUrl, errs)
		return
	}
	enc := determineEncoding(body)
	if enc != "utf-8" {
		body = decoderConvert(enc, body)
	}

	level := c.Urls[targetUrl]
	for _, keyword := range c.Keywords {
		if strings.Contains(body, keyword) {
			fmt.Println(keyword, "is contained.")
			level = 1
		}
	}
	doc := soup.HTMLParse(body)
	links := doc.FindAll("a")
	u, _ := url.Parse(targetUrl)

	for _, a := range links {
		link := a.Attrs()["href"]
		if strings.HasPrefix(link, "http") {
			c.AddInUrls(link, level)
		} else if strings.HasPrefix(link, "/") {
			c.AddInUrls(u.Scheme+"://"+u.Host+link, level)
		}
	}
	r := regexp.MustCompile("(.*/)")

	path := "files/" + u.Host + u.Path

	if len(u.Path) == 0 {
		os.MkdirAll(path, 0666)

	} else {
		ps := r.FindStringSubmatch(path)
		log.Println(targetUrl)
		os.MkdirAll(ps[0], 0666)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	var title, content string
	htmlBody := doc.Find("body")
	htmlTitle := doc.Find("title")
	if htmlBody.Error == nil {
		content = htmlBody.FullText()
	}
	if htmlTitle.Error == nil {
		title = htmlTitle.Text()
	}
	c.VisitedUrls[targetUrl] = struct{}{}
	p := &Page{
		Url:       targetUrl,
		Title:     title,
		Level:     level,
		ParentUrl: parentUrl,
		TaskID:    taskID,
	}
	db.Save(p)
	f.WriteString(content)
}

func (c *Collector) AddInUrls(link string, level int) {
	l, ok := c.Urls[link]
	if !ok {
		c.Urls[link] = level + 1
	} else {
		if l > level+1 {
			c.Urls[link] = level + 1
		}
	}
}
