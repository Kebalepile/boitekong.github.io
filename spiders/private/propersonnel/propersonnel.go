package propersonnel

import (
	"context"
	"fmt"
	"github.com/Kebalepile/job_board/spiders/types"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"sync"
	"time"
)

// ProPesonnel package spider type
type Spider struct {
	Name           string
	AllowedDomains []string
	Shutdown       context.CancelFunc
	Posts          types.ProPersonnelJobs
}

// initiate the Spider instant
// Configers chromedp options such as headless flag userAgent & window size
// Creates Navigates to the allowed domain to crawl
func (s *Spider) Launch(wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println(s.Name, " spider has Lunched ", s.Date())
	s.Posts.Title = fmt.Sprintf("propersonnel-jobs%s", s.Date())

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // set headless to true for production
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.WindowSize(768, 1024), // Tablet size
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)

	s.Shutdown = cancel

	log.Println("Loading ", s.Name)

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(s.AllowedDomains[0]))
	s.Error(err)

	s.Robala(5)

	menuSelector := `.mmenu-trigger`
	var nodes []*cdp.Node

	err = chromedp.Run(
		ctx,
		chromedp.WaitVisible(menuSelector),
		chromedp.Click(menuSelector),
		chromedp.WaitVisible(`#mobile-nav`),
		chromedp.Nodes(`#mobile-nav a`, &nodes, chromedp.ByQueryAll))
	s.Error(err)

	if n := len(nodes); n > 0 {
		var href string
		for _, n := range nodes {
			href = n.AttributeValue("href")
			if yes := strings.Contains(href, "vacancies"); yes {
				break
			}
		}
		s.Robala(5)
		s.vacancies(href, ctx)
	}
}

// scrapes availabe job posts on loaded page url
// adds them to Posts.Links slice
// once done save the information to a *.json file
func (s *Spider) vacancies(url string, ctx context.Context) {

	log.Println("Loading vacancies page")

	var t string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Title(&t))

	s.Error(err)
	s.Robala(10)

	if yes := strings.Contains(strings.ToLower(t), "vacancies"); yes {
		selector := `#advert_list`
		err = chromedp.Run(ctx,
			chromedp.WaitEnabled(selector),
			chromedp.ScrollIntoView(selector))
		s.Error(err)
	}
}

func (s *Spider) Date() string {
	t := time.Now()
	return t.Format("02 January 2006")
}

// closes chromedp broswer instance
func (s *Spider) Close() {
	log.Println(s.Name, "is done.")
	s.Shutdown()
}
func (s *Spider) Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// pauses spider for given duration
func (s *Spider) Robala(second int) {
	time.Sleep(time.Duration(second) * time.Second)
}
