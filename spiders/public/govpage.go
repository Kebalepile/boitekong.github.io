package public

import (
	"context"
	"github.com/Kebalepile/job_board/spiders"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

type Spider struct {
	Name           string
	AllowedDomains []string
	Shutdown       context.CancelFunc
}

func (s *Spider) Launch() {

	log.Println(s.Name, " spider has Lunched ", s.Date())

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // set headless false
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.WindowSize(768, 1024), // Tablet size
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)

	s.Shutdown = cancel

	err := chromedp.Run(ctx,
		chromedp.Navigate(s.AllowedDomains[0]),
	)
	s.Error(err)

	var nodes []*cdp.Node

	err = chromedp.Run(ctx,
		chromedp.Click(`*[aria-label='Menu']`),
		chromedp.Nodes(`ul li.wsite-menu-item-wrap a.wsite-menu-item`, &nodes, chromedp.ByQueryAll))
	s.Error(err)

	// loop over the anchor elements
	for _, n := range nodes {
		var (
			text string
			url  string
		)
		err = chromedp.Run(ctx,
			chromedp.TextContent(n.FullXPath(), &text),
			chromedp.Location(&url))
		s.Error(err)
		if match := strings.Contains(strings.ToLower(text), "govpage"); match {
			href := n.AttributeValue("href")
			// remove first '/' from href as url ends with '/'
			govUpdates := url + href[1:]

			selector := `.blog-title-link`

			log.Println("Loading government updates page")

			err = chromedp.Run(ctx,
				chromedp.Navigate(govUpdates),
				chromedp.WaitEnabled(selector),
				chromedp.ScrollIntoView(selector),
				chromedp.Location(&url))

			s.Error(err)
			if n := strings.Compare(url, s.AllowedDomains[1]); n == 0 {
				s.vacancies(ctx, selector)
			}
			break
		}

	}
	s.Close()
}

func (s *Spider) vacancies(ctx context.Context, selector string) {
	log.Println("Searching for latest government vacancies.")

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Nodes(selector, &nodes, chromedp.ByQueryAll))
	s.Error(err)

	for _, n := range nodes {
		var text string
		err = chromedp.Run(ctx,
			chromedp.TextContent(n.FullXPath(), &text))
		s.Error(err)
		title := strings.ToLower(text)
		if match := strings.Contains(title, strings.ToLower(s.Date())); match {
			govpageLinks := spiders.Links{
				Title: title,
			}
			href := n.AttributeValue("href")
			url := "https://" + href[2:]

			s.links(ctx, url, govpageLinks)

		} else {
			log.Println("Sorry, No Government Job Posts for today")
		}
	}

}
func (s *Spider) links(ctx context.Context, url string, govpageLinks spiders.Links) {

	log.Println("Searching For Advert Post Links")

	selector := `[id^='blog-post-'] a`

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitEnabled(selector, chromedp.ByQueryAll),
		chromedp.ScrollIntoView(selector),
		chromedp.Nodes(selector, &nodes, chromedp.ByQueryAll))

	s.Error(err)

	log.Println("Found some posts, number of posts deteched is: ", len(nodes))

	if posts := len(nodes); posts > 0 {

		for _, n := range nodes {

			var text string
			href := n.AttributeValue("href")

			govpageLinks.Links[text] = href

			err = chromedp.Run(ctx,
				chromedp.TextContent(n.FullXPath(), &text))
			s.Error(err)
		}
		log.Println(govpageLinks)
		
	}
}
func (s *Spider) Date() string {
	t := time.Now()
	return t.Format("02 January 2006")
}
func (s *Spider) Close() {
	log.Println(s.Name, "is done.")
	s.Shutdown()

}
func (s *Spider) Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
