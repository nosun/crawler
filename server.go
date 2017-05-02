package main

import (
	"hash/fnv"
	"net/url"
	spider "github.com/sundy-li/wechat_spider"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

/**
	## mysql server in localhost:4000
	create database wechat;
	use wechat;
	create table wx_article(id bigint(15) not null primary key , url varchar(512), data text, update_time bigint(15));
**/

func main() {
	var port = "8899"
	spider.InitConfig(&spider.Config{
		Verbose: true,    // 日志查看
		AutoScroll: true, // 自动翻页
		SleepSecond: 2,   // 间歇时间
	})
	spider.Regist(&CustomProcessor{})
	spider.Run(port)
}

//Just to implement Output Method of interface{} Processor
type CustomProcessor struct {
	spider.BaseProcessor
}

type M map[string]interface{}

func (c *CustomProcessor) Output() {
	for _, result := range c.UrlResults() {

		uri, _ := url.ParseRequestURI(result.Url)

		biz := uri.Query().Get("__biz")
		mid := uri.Query().Get("mid")
		idx := uri.Query().Get("idx")
		aid := hash(biz + "_" + mid + idx)

		stmt, err := db.Prepare(
			"insert ignore into wx_article(aid,biz,mid,idx,url,status) values(?,?)")

		if err != nil {
			println(err.Error())
			continue
		}

		res, err := stmt.Exec(aid, biz, mid, idx, result.Url, 0)

		id, err := res.LastInsertId()
		checkErr(err)

		stmt, err = db.Prepare("insert into wx_article_data (id,content) values (?)")
		checkErr(err)

		_, err = stmt.Exec(id)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wechat?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func hash(s string) int64 {
	h := fnv.New32()
	h.Write([]byte(s))
	return int64(h.Sum32())
}

var (
	db *sql.DB
)
