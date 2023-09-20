package main

import(
	"log"
	 "github.com/Kebalepile/job_board/spiders/public/govpage"
)
func main(){
	log.Println("Job Board Scrapper Initiated ")
	govpageSpider := govpage.Spider{
		Name: "gov-page",
		AllowedDomains: []string{"https://www.govpage.co.za/",
		"https://www.govpage.co.za/latest-govpage-updates"},
	}

	govpageSpider.Launch()
}