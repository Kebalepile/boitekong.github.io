<<<<<<< HEAD
package propersonnel

import (
	"context"
	"fmt"
	"github.com/Kebalepile/job_board/pipeline"
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
		chromedp.Flag("headless", true), // set headless to true for production
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.WindowSize(768, 1024), // Tablet size
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)

	s.Shutdown = cancel

	log.Println("Loading ", s.Name)

	menuSelector := `.mmenu-trigger`
	var nodes []*cdp.Node

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(s.AllowedDomains[0]),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(menuSelector),
		chromedp.Click(menuSelector),
		chromedp.Sleep(5*time.Second),
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
		chromedp.Sleep(5*time.Second),
		chromedp.Title(&t))

	s.Error(err)

	if yes := strings.Contains(strings.ToLower(t), "vacancies"); yes {

		iframe := `#advanced_iframe`

		var (
			nodes []*cdp.Node
			url   string
		)

		err = chromedp.Run(ctx,
			chromedp.Nodes(iframe, &nodes, chromedp.ByQuery))
		s.Error(err)

		if len(nodes) > 0 {

			url = nodes[0].AttributeValue("src")
			selector := `#advert_list`

			err = chromedp.Run(ctx,
				chromedp.Navigate(url),
				chromedp.Sleep(5*time.Second),
				chromedp.WaitVisible(selector, chromedp.ByQuery),
				chromedp.ScrollIntoView(selector),
				chromedp.Nodes(`.job-spec`, &nodes, chromedp.ByQueryAll))
			s.Error(err)

			log.Println(len(nodes), " job posts found on ", s.Name)

			if len(nodes) > 0 {
				for _, n := range nodes {
					id := n.AttributeValue("id")
					expression := fmt.Sprintf(`(() => {    

							const html = document.querySelector("#%s").innerHTML;

							let regexPattern = /<div\s+class="job-spec-value">(.*?)<\/div>/;
							let match = html.match(regexPattern);
							const jobTitle = match ? match[1].trim() : "";

							regexPattern = /<div[^>]*id="start_date"[^>]*>([\s\S]*?<div\s+class="job-spec-value">(.*?)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const startDate = match ? match[2].trim() : "";

							regexPattern = /<div[^>]*id="job_spec_type"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bjob-spec-value\b[^"]*"[^>]*>([^<]*)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const vacancyType = match ? match[2].trim() : "";

							regexPattern = /<div[^>]*id="sectors_list"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bbadge-ribbon\b[^"]*"[^>]*><i[^>]*><\/i>\s*([^<]*)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const jobSpecFields = match ? match[2].trim() : "";
					
							regexPattern = /<span\s+id="region_label">(.*?)<\/span>/; // Use the regex pattern to match and extract the date content
							match = html.match(regexPattern);
							const location = {
							region: match ? match[1].trim() : "",
							};
					
							regexPattern = /<span\s+id="location_label">(.*?)<\/span>/; // Use the regex pattern to match and extract the date content
							match = html.match(regexPattern);
					
							location.city = match ? match[1].trim() : "";
					
							regexPattern = /<div[^>]*id="contact"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bjob-spec-value\b[^"]*"[^>]*>([^<]*)<\/div>[\s\S]*?)<\/div>/; // Use the regex pattern to match and extract the date content
							match = html.match(regexPattern);
							const contact = match ? match[2].trim() : "";
					
							regexPattern =
							/<div[^>]*id="description"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bjob-spec-value\b[^"]*"[^>]*>([\s\S]*?)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const details = match ? match[2].trim().replace(/<[^>]*>/g, "") : "";
						
							regexPattern =
							/<div[^>]*id="apply_button"[^>]*>([\s\S]*?<a[^>]*href="([^"]*)"[^>]*>[\s\S]*?<\/a>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const apply = window.location.origin + (match ? match[2].trim() : "");
					
							return {
								jobTitle,
								startDate,
								vacancyType,
								jobSpecFields,
								location,
								details,
								contact,
								apply,
							};
					})()`, id)

					var (
						JobPost types.ProJobPost
					)
					err = chromedp.Run(ctx,
						chromedp.ScrollIntoView(id),
						chromedp.Evaluate(expression, &JobPost))
					s.Error(err)

					s.Posts.Links = append(s.Posts.Links, JobPost)

				}

				err = pipeline.ProPersonnelFile(&s.Posts)
				if err != nil {
					s.Error(err)
				}
				s.Close()
			} else {
				s.Close()
			}
		}

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
=======
package propersonnel

import (
	"context"
	"fmt"
	"github.com/Kebalepile/job_board/pipeline"
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
		chromedp.Flag("headless", true), // set headless to true for production
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.WindowSize(768, 1024), // Tablet size
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)

	s.Shutdown = cancel

	log.Println("Loading ", s.Name)

	menuSelector := `.mmenu-trigger`
	var nodes []*cdp.Node

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(s.AllowedDomains[0]),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(menuSelector),
		chromedp.Click(menuSelector),
		chromedp.Sleep(5*time.Second),
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
		chromedp.Sleep(5*time.Second),
		chromedp.Title(&t))

	s.Error(err)

	if yes := strings.Contains(strings.ToLower(t), "vacancies"); yes {

		iframe := `#advanced_iframe`

		var (
			nodes []*cdp.Node
			url   string
		)

		err = chromedp.Run(ctx,
			chromedp.Nodes(iframe, &nodes, chromedp.ByQuery))
		s.Error(err)

		if len(nodes) > 0 {

			url = nodes[0].AttributeValue("src")
			selector := `#advert_list`

			err = chromedp.Run(ctx,
				chromedp.Navigate(url),
				chromedp.Sleep(5*time.Second),
				chromedp.WaitVisible(selector, chromedp.ByQuery),
				chromedp.ScrollIntoView(selector),
				chromedp.Nodes(`.job-spec`, &nodes, chromedp.ByQueryAll))
			s.Error(err)

			log.Println(len(nodes), " job posts found on ", s.Name)

			if len(nodes) > 0 {
				for _, n := range nodes {
					id := n.AttributeValue("id")
					expression := fmt.Sprintf(`(() => {    

							const html = document.querySelector("#%s").innerHTML;

							let regexPattern = /<div\s+class="job-spec-value">(.*?)<\/div>/;
							let match = html.match(regexPattern);
							const jobTitle = match ? match[1].trim() : "";

							regexPattern = /<div[^>]*id="start_date"[^>]*>([\s\S]*?<div\s+class="job-spec-value">(.*?)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const startDate = match ? match[2].trim() : "";

							regexPattern = /<div[^>]*id="job_spec_type"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bjob-spec-value\b[^"]*"[^>]*>([^<]*)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const vacancyType = match ? match[2].trim() : "";

							regexPattern = /<div[^>]*id="sectors_list"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bbadge-ribbon\b[^"]*"[^>]*><i[^>]*><\/i>\s*([^<]*)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const jobSpecFields = match ? match[2].trim() : "";
					
							regexPattern = /<span\s+id="region_label">(.*?)<\/span>/; // Use the regex pattern to match and extract the date content
							match = html.match(regexPattern);
							const location = {
							region: match ? match[1].trim() : "",
							};
					
							regexPattern = /<span\s+id="location_label">(.*?)<\/span>/; // Use the regex pattern to match and extract the date content
							match = html.match(regexPattern);
					
							location.city = match ? match[1].trim() : "";
					
							regexPattern = /<div[^>]*id="contact"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bjob-spec-value\b[^"]*"[^>]*>([^<]*)<\/div>[\s\S]*?)<\/div>/; // Use the regex pattern to match and extract the date content
							match = html.match(regexPattern);
							const contact = match ? match[2].trim() : "";
					
							regexPattern =
							/<div[^>]*id="description"[^>]*>([\s\S]*?<div[^>]*class="[^"]*\bjob-spec-value\b[^"]*"[^>]*>([\s\S]*?)<\/div>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const details = match ? match[2].trim().replace(/<[^>]*>/g, "") : "";
						
							regexPattern =
							/<div[^>]*id="apply_button"[^>]*>([\s\S]*?<a[^>]*href="([^"]*)"[^>]*>[\s\S]*?<\/a>[\s\S]*?)<\/div>/;
							match = html.match(regexPattern);
							const apply = window.location.origin + (match ? match[2].trim() : "");
					
							return {
								jobTitle,
								startDate,
								vacancyType,
								jobSpecFields,
								location,
								details,
								contact,
								apply,
							};
					})()`, id)

					var (
						JobPost types.ProJobPost
					)
					err = chromedp.Run(ctx,
						chromedp.ScrollIntoView(id),
						chromedp.Evaluate(expression, &JobPost))
					s.Error(err)

					s.Posts.Links = append(s.Posts.Links, JobPost)

				}

				err = pipeline.ProPersonnelFile(&s.Posts)
				if err != nil {
					s.Error(err)
				}
				s.Close()
			} else {
				s.Close()
			}
		}

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
>>>>>>> 3c99bb87c76a5385a5e756e6b1a05acdd46aaf36
