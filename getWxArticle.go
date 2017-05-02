package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getArticle(id int,wxUrl string)  {

	resp, err := http.Get(wxUrl)
	checkErr(err)

	doc,err := goquery.NewDocumentFromReader(resp.Body)
	// strings.NewReader(data) // 从字符串中读取
	checkErr(err)

	origin,_ := doc.Find("#copyright_logo").First().Html()
	author,_ := doc.Find("#post-user").First().Html()
	title,_  := doc.Find("#activity-name").First().Html()
	date,_   := doc.Find("#post-date").First().Html()
	content,_ := doc.Find("#js_content").First().Html()

	stmt, err := db.Prepare("update wx_article set " +
		"title = ? , author = ?, date = ?, origin = ? , status = 1 where id = ?)")
	checkErr(err)

	_, err = stmt.Exec(strings.TrimSpace(title), strings.TrimSpace(author), strings.TrimSpace(date), origin, id)
	checkErr(err)

	stmt, err = db.Prepare("update wx_article_data set content = ?  where id = ?")
	checkErr(err)

	_, err = stmt.Exec(content, id)
	checkErr(err)
}

func main() {

	rows, err := db.Query("select id, url from wx_article where status = 0")
	checkErr(err)

	for rows.Next() {
		var id int
		var wxUrl string
		rows.Columns()
		err = rows.Scan(&id, &wxUrl)
		checkErr(err)
		fmt.Println(time.Now(),wxUrl)
		getArticle(id, wxUrl)
		time.Sleep(time.Second * 1)
	}
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wechat?timeout=100m&charset=utf8")
	if err != nil {
		panic(err)
	}
}

var (
	db *sql.DB
)
