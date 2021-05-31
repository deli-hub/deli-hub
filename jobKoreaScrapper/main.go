package main

import (
	"os"
	"strings"

	"github.com/deli-hub/jobKoreaScrapper/practice"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(practice.CleanString(c.FormValue("term")))
	practice.Scrape(term)

	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	// 루트와 e.Start는 함께 있어야 에코 작동 가능
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/getJobs", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
