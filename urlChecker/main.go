package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request Failed")

func main() {
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
		// hit : 인터넷 웹 서버의 파일 하나에 접속하는 것을 뜻함
		hitUrl(url)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking:", url)
	// 아래는 Go의 standard library
	resp, err := http.Get(url)
	// url이 정상적인지 확인한 후에 에러가 나면 return
	// StatusCode는 0부터 300까지는 리다이렉션을 해주고 400부터는 에러가 발생한다.
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
