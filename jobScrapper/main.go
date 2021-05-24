package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=java&limit=50"

// goquery document
// https://github.com/puerkitobio/goquery

func main() {
	// 여기서의 jobs는 많은 배열의 모임이다.
	var jobs []extractedJob
	totalPages := getPages()
	// fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		// 총 페이지 수만큼 for문을 돌림
		extractedJobs := getPage(i)
		// what we're going to do is to add contents of the extractedJobs and make it one
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted: ", len(jobs))

}

func getPage(page int) []extractedJob {
	var jobs []extractedJob // job is a slice of extractedJob
	// strconv.Itoa (int -> string) : Go 내장함수
	pageUrl := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting pageURL:: ", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// 페이지 내에서 공고란 div
	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		// jobs에 추출해낸 job을 append한다.
		jobs = append(jobs, job)
	})

	return jobs
}

func getPages() int {
	pages := 0
	// baseURL에 대한 res, err를 반환
	res, err := http.Get(baseURL)
	if err != nil {
		checkErr(err)
	}
	if res.StatusCode != 200 {
		checkCode(res)
	}

	// res.Body는 기본적으로 byte이며 IO이다.
	// 따라서 getPages() 함수가 끝났을 때 response를 닫아주어야한다..?
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// Find the review items
	// html에서 클래스명 = pagination 찾음
	// each의 사용법은 doc에 있다.
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// a 태그를 찾아서 한 페이지당 몇 개의 링크를 포함하고 있는지 확인
		pages = s.Find("a").Length()
	})

	return pages
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

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title > a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}

}

// 모든 공백을 제거한다.
func cleanString(str string) string {
	// Join concatenates the elements of its first argument to create a single string
	// Fields splits the string s around each instance of one or more consecutive white space characters
	// TrimSpace returns a slice of the string s, with all leading and trailing white space removed
	/*
		예를 들어 Fields가 "Hello       world        !"를 "Hello""world""!"로 모아준다면,
		Join은 "Hello world !"로 스페이스와 함께 하나의 string로서 반환해준다.
	*/
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

// job scrapper 데이터를 csv로 변환
// https://golang.org/pkg/encoding/csv/#example_Writer
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	// 함수가 끝나는 시점에 파일에 데이터를 입력 (defer : 함수가 끝나면 호출하는 함수)
	// Flush writes any buffered data to the underlying io.Writer. To check if an error occurred during the Flush, call Error.
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}
