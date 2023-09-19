package main

import(
	"log"
	pub "github.com/Kebalepile/job_board/spiders/public"
)
func main(){
	log.Println("Job Board Scrapper Initiated ")
	govpageSpider := pub.Spider{
		Name: "gov-page",
		AllowedDomains: []string{"https://www.govpage.co.za/",
		"https://www.govpage.co.za/latest-govpage-updates"},
	}

	govpageSpider.Launch()
}