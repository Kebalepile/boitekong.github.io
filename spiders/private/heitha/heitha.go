package heitha

import(
	"context"
	// "github.com/Kebalepile/job_board/spiders"
	// "github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	// "strings"
	// "time"
)
// phase 1
// 1. go to http://www.heitha.co.za/
// 2 . click button with class .all-button and textContent
// 3. of 'See all jobs > ' it will navigate to http://www.heitha.co.za/jobs
// 4. get div of .col-sm-9.items div innerHTML
// 5. get each a tag href and get it's child .row.job > .col-sm-8
// 6. get the industy title textContent .industry-title
// 7. get the job title textContent .job-title
// 8. get the bullets textContent .bullets
// 9. get the Expiry date textContent  .col-sm-4  > .expiry-date (2 of them)
// 10. scroll page navigation into view here is the xpath //*[@id="layout-content"]/div[2]/div/div[2]/div
// 11. get the pagination ul .pagination & look for li iwth class of .active check if it's next sibling has clss of disabled
// 12. if not click the next sibling and start form 1 to 11 for the whole page
// broswer just navigated to.
// do this up to page 10 if the url contains page 10 http://www.heitha.co.za/jobs?page=10
//  an the active li.active > span textContent contains 10 stop phase 1.

// phase 2
// 1. for the a tag href retrived in phase 1 section 5 naviage to the href
// 2. get job content div with class ".col-sm-8.job-content"
// 3. get the latter divs innerHTML
// 4. save along side phase 1 content
// 5. do this for all other a tags
