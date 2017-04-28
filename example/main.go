package main

import "net/http"
import (
	"github.com/nosun/crawler"
	"github.com/nosun/crawler/tools/agent"
)

func main() {

	task := "http://www.cwzg.cn"
	client := &http.Client{}

	spider := &crawler.Spider{
		UserAgent: agent.GetRandomUserAgent(),
		Client: client,
	}

	spider.Run(task,"GET")
}
