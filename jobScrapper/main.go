package main

import (
	"os"
	"strings"

	"github.com/deli-hub/jobScrapper/scrapper"
	"github.com/labstack/echo"
)

const File_Name string = "jobs.csv"

// echo => https://echo.labstack.com/guide
func handleHome(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	// 사용자가 파일을 다운로드하면 서버에서는 해당 파일을 지운다.
	defer os.Remove(File_Name)
	// FormValue -> form안의 요소 중 괄호에 해당하는 값을 가져온다.
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	// 첨부 파일을 리턴해준다. (다운받을 파일, 사용자에게 전달할 때 사용할 이름)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
