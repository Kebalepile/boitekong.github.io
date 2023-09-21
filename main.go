package main

import (
	"github.com/Kebalepile/job_board/spiders/private/heitha"
	"github.com/Kebalepile/job_board/spiders/public/govpage"
	"github.com/Kebalepile/job_board/spiders/types"
	"log"
	"sync"
)



func main() {
	log.Println("Job Board Scrapper Initiated ")
	govpageSpider := govpage.Spider{
		Name: "gov-page",
		AllowedDomains: []string{
			"https://www.govpage.co.za/",
			"https://www.govpage.co.za/latest-govpage-updates",
		},
	}

	heithaSpider := heitha.Spider{
		Name: "heitha",
		AllowedDomains: []string{
			"http://www.heitha.co.za/",
			"http://www.heitha.co.za/jobs",
		},
	}

	goFuncs := []types.Crawler{
		&govpageSpider,
		&heithSpider
	}

	var wg sync.WaitGroup
	for _, f := range goFuncs {
		wg.Add(1)
		go f.Launch(&wg)
	}
	wg.Wait()
}
