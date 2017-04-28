package crawler

import (
	"net/http"
	"log"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"net/url"
	"io/ioutil"
)

// Doer defines the method required to use a type as HttpClient.
// The net/*http.Client type satisfies this interface.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Spider struct {
	UserAgent string
	Client Doer
}

//以Must前缀的方法或函数都是必须保证一定能执行成功的,否则将引发一次panic
var atagRegExp = regexp.MustCompile(`<a[^>]+[(href)|(HREF)]\s*\t*\n*=\s*\t*\n*[(".+")|('.+')][^>]*>[^<]*</a>`)


// process
func (s *Spider) Run (taskUrl string, method string) {

	req,err := http.NewRequest(method,taskUrl,nil)
	if err != nil {
		fmt.Println("request error")
		return
	}

	resp,err := s.Client.Do(req)

	if err != nil {
		fmt.Println("send request error")
		return
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("request status code error:", resp.StatusCode, string(body) )
	}

	if err != nil {
		fmt.Println("read body error")
		return
	}

	// analysis
	var urls [] string

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")

		// Resolve address
		u, err := url.Parse(val)

		//fmt.Println(u)
		if err != nil {
			fmt.Println("parase url error")
			return
		}

		if res,err := regexp.MatchString("^http(.+)html$",u.String()); err == nil && res == true {
			urls = append(urls,u.String())
		}

		for i,v := range urls {
			fmt.Println(i,v)
		}
	})

	// handle output



}

