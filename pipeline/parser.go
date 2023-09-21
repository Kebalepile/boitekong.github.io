package pipeline

import (
	"golang.org/x/net/html"
	"log"
	"strings"
)

type HtmlParser struct {
	Html string
}

func (p *HtmlParser) Init() {
	log.Println(p.Html)
	doc, err := html.Parse(strings.NewReader(p.Html))
	if err != nil {
		panic(err)

	}
	// Call a recursive function to extract the desired information
	p.extractInfo(doc)
}

func (p *HtmlParser) extractInfo(n *html.Node) {
	log.Println(n == nil)
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		log.Println(n.Data)
		switch n.Data {
		case "a":
			// Check if the <a> tag has an href attribute
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					log.Println("href: ", attr.Val)
					break
				}
			}

			// Find .row.job > .col-sm-8
			var industryTitle, jobTitle, bullets string
			var expiryDates []string

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "div" {
					class := p.getAttributeValue(c, "class")

					switch class {
					case "industry-title":
						industryTitle = p.getTextContent(c)
					case "job-title":
						jobTitle = p.getTextContent(c)
					case "bullets":
						bullets = p.getTextContent(c)
					}
				}
			}
			log.Println("Industry Title:", industryTitle)
			log.Println("Job Title:", jobTitle)
			log.Println("Bullets:", bullets)

			// Find Expiry Dates
			for c := n.LastChild; c != nil; c = c.PrevSibling {
				if c.Type == html.ElementNode && c.Data == "div" {
					class := p.getAttributeValue(c, "class")
					if class == "expiry-date" {
						expiryDates = append(expiryDates, p.getTextContent(c))
						if len(expiryDates) == 2 {
							break
						}
					}
				}
			}

			log.Println("Expiry Dates:", expiryDates)

		default:
			// Recursively process child nodes
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				p.extractInfo(c)
			}
		}
	}

}

func (p *HtmlParser) getAttributeValue(n *html.Node, attrName string) string {
	for _, attr := range n.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}

func (p *HtmlParser) getTextContent(n *html.Node) string {
	var textContent string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			textContent += c.Data
		}
	}
	return strings.TrimSpace(textContent)
}
