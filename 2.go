package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {

		url := "http://pvp.qq.com/web201605/herodetail/106.shtml"
		doc,_ := goquery.NewDocument(url)
		ele := doc.Find("div.show-list")
		ele.Each(func(i int, content *goquery.Selection) {
			name := content.Find("p.skill-name").Find("b").Eq(0).Text()
			fmt.Println(i,name)
		})
	}

