package practice

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
	gno      string
	title    string
	exp      string
	edu      string
	location string
	date     string
}

func Scrape(term string) {
	var baseURL string = "https://www.jobkorea.co.kr/Search/?stext=" + term
	var jobs []extractedJob

	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalPages; i++ {
		// extractedJobs := getPage(i, baseURL)
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted: ", len(jobs))
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)

	if err != nil {
		checkErr(err)
	}
	if res.StatusCode != 200 {
		checkCode(res)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".recruit-info").Each(func(i int, s *goquery.Selection) {
		pages = s.Find(".wide").Find("a").Length()
	})
	return pages
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	// 채널에서 호환하는 객체 타입
	c := make(chan extractedJob)

	pageUrl := url + "&tabType=recruit&Page_No=" + strconv.Itoa(page)
	fmt.Println("Requesting pageURL:: ", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".list-post")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		// job := extractJob(card)
		//jobs = append(jobs, job)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		if job.gno != "" {
			jobs = append(jobs, job)
		}
	}

	// return jobs
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	gno, _ := card.Attr("data-gno")
	titleArea := card.Find(".post-list-info > .title")
	title, _ := titleArea.Attr("title")
	exp := CleanString(card.Find(".option > .exp").Text())
	edu := CleanString(card.Find(".option > .edu").Text())
	location := CleanString(card.Find(".option > .loc.long").Text())
	date := CleanString(card.Find(".option > .date").Text())

	//return extractedJob{
	c <- extractedJob{
		gno:      gno,
		title:    title,
		exp:      exp,
		edu:      edu,
		location: location,
		date:     date,
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	defer w.Flush()

	headers := []string{"gno", "Title", "Experience", "Education", "Location", "Due-date"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		if job.gno != "" {
			jobSlice := []string{"https://www.jobkorea.co.kr/Recruit/GI_Read/" + job.gno + "?Oem_Code=C1&logpath=1", job.title, job.exp, job.edu, job.location, job.date}
			jwErr := w.Write(jobSlice)
			checkErr(jwErr)
		}
	}

}

// 에러체크
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// response에서 StatusCode를 확인
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
