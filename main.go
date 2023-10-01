package main

import (
	// "github.com/Kebalepile/job_board/spiders/private/heitha"
	// "github.com/Kebalepile/job_board/spiders/public/govpage"
	"github.com/Kebalepile/job_board/spiders/private/propersonnel"
	"github.com/Kebalepile/job_board/spiders/types"
	"log"
	"sync"
)

func main() {
	log.Println("Job Board Scrapper Initiated ")
	// govpageSpider := govpage.Spider{
	// 	Name: "gov-page",
	// 	AllowedDomains: []string{
	// 		"https://www.govpage.co.za/",
	// 		"https://www.govpage.co.za/latest-govpage-updates",
	// 	},
	// }

	// heithaSpider := heitha.Spider{
	// 	Name: "heitha-page",
	// 	AllowedDomains: []string{
	// 		"http://www.heitha.co.za/",
	// 		"http://www.heitha.co.za/jobs",
	// 	},
	// }

	propersonnelSpider := propersonnel.Spider{
		Name:"Pro-Personnel",
		AllowedDomains: []string{
			"https://www.pro-personnel.co.za/",
			"https://www.pro-personnel.co.za/vacancies/",
		},
	}

	goFuncs := []types.Crawler{
		// &govpageSpider,
		// &heithaSpider,
		&propersonnelSpider,
	}

	var wg sync.WaitGroup
	for _, f := range goFuncs {
		wg.Add(1)
		go f.Launch(&wg)
	}
	wg.Wait()
}
