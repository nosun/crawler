package main

import (
	"regexp"
)

func main() {
	url := "http://www.cwzg.cn/politics/201611/32600.html"

	res,_ := regexp.MatchString("^http(.+)\\[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",url)

	println(res)
}
