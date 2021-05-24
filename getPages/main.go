package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=java&limit=50"

// goquery document
// https://github.com/puerkitobio/goquery

func main() {
	pages := getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	if err != nil {
		checkErr(err)
	}
	if res.StatusCode != 200 {
		checkCode(res)
	}

	defer res.Body.Close()

	// res.Body는 기본적으로 byte이며 IO이다.
	// 따라서 getPages() 함수가 끝났을 때 response를 닫아주어야한다..?
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// Find the review items
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
