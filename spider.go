package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func skill_get(url string) {
	body:=trans(url)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println(err.Error())
	}
	ele := doc.Find("div.show-list")
	ele.Each(func(i int, content *goquery.Selection) {
		name := content.Find("p.skill-name").Find("b").Eq(0).Text()
		fmt.Println(i,name)
	})
}
func name_get(url string) {
	body:=trans(url)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println(err.Error())
	}
	ele := doc.Find("div.cover")
	ele.Each(func(i int, content *goquery.Selection) {
		name := content.Find("h2.cover-name").Eq(0).Text()
		fmt.Println(name)
	})
}
func title_get(url string) {
	body:=trans(url)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println(err.Error())
	}
	ele := doc.Find("div.cover")
	ele.Each(func(i int, content *goquery.Selection) {
		name := content.Find("h3.cover-title").Eq(0).Text()
		fmt.Println(name)
	})
}
func trans(url string)(string){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	by, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	body := string(by)
	body, err = iconv.ConvertString(body, "GBK", "utf-8")
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
func main()  {
	var num int
	for num=100;num<600 ;num++  {
		url := "http://pvp.qq.com/web201605/herodetail/"+strconv.Itoa(num)+".shtml"
		name_get(url)
		title_get(url)
		skill_get(url)
	}

}