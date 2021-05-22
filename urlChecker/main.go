package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request Failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	// var results = map[string]string{}
	// make(map) => map을 만들어주는 함수(초기화까지 해줌)
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	}

	for _, url := range urls {
		// result := "OK"
		// hit : 인터넷 웹 서버의 파일 하나에 접속하는 것을 뜻함
		// err :=
		go hitUrl(url, c)
		// if err != nil {
		// result = "FAILED"
		// }
		// 해당 url을 key로, err여부(result)를 value로
		// results[url] = result
	}

	// 터미널에 찍히는 값을 편하게 보기 위함
	// for _, result := range results {
	// 	fmt.Println(result)
	// }

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

/* CHECKER 예제 함수*/
/*
func hitUrl(url string) error {
	fmt.Println("Checking:", url)
	// 아래는 Go의 standard library
	resp, err := http.Get(url)
	// url이 정상적인지 확인한 후에 에러가 나면 return
	// StatusCode는 0부터 300까지는 리다이렉션을 해주고 400부터는 에러가 발생한다.
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}*/

/* CHANNEL 예제 함수*/
// 이 함수는 채널이 보내기만 가능하다 👇
func hitUrl(url string, c chan<- requestResult) {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
