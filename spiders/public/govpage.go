package public

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	// "strings"
)

type Spider struct {
	Name           string
	AllowedDomains []string
}

func (s *Spider) Launch() {
	log.Println(s.Name, " spider has Lunched !")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // set headless false
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.WindowSize(768, 1024), // Tablet size
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)

	err := chromedp.Run(ctx,
		chromedp.Navigate(s.AllowedDomains[0]),
	)
	s.error(err)

	var nodes []*cdp.Node

	err = chromedp.Run(ctx,

		chromedp.WaitReady(`body`),

		chromedp.Click(`*[aria-label='Menu']`),

		chromedp.Nodes(`ul li.wsite-menu-item-wrap a.wsite-menu-item`, &nodes, chromedp.ByQueryAll))
	s.error(err)
	// loop over the anchor elements
	for _, n := range nodes {
		// var textContent string
		log.Println("------------------TOP---------------")
		log.Println(&n)
		log.Println(n)
		log.Println("------------------BOTTOM---------------")
	
		
	}
	s.done()
}

func (s *Spider) vacancies(ctx context.Context){
	log.Println("Searching for latest government vacancies.")
	// check if url include s.AllowedDomains[1]
	// get current date in day month year format
	// wait for .blog-title-link to load
	// get elements ".blog-title-link"
	// run this code
	/**
	   const targetHandle = await elements.reduce(
            async (targetHandle, elementHandle) => {
              const textContent = await page.evaluate(
                (elem) => elem.textContent,
                elementHandle
              );
              if (textContent.includes(currentDate)) {
                targetHandle = elementHandle;
                return targetHandle;
              }
              return targetHandle;
            },
            null
          );

          if (targetHandle) {
            await targetHandle?.click();

            await this.#advertLinks(page);
          }else {
            fs.writeFile(
              this.#databasePath(this.#date("date")),
              JSON.stringify(
                {
                  text: "No job posts for today",
                  date: this.#date("date"),
                },
                null,
                4
              ),
              (error) =>
                error
                  ? console.log(error.message)
                  : console.log(`${this.#date("date")}.json save to database`)
            );
            await this.#terminate();
		  */

}
func (s *Spider ) links(ctx context.Context){

}
func (s *Spider) done() {
	log.Println(s.Name, "is done.")
}
func (s *Spider) error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
