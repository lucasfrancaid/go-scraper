package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Shopify struct {
	URL       string
	KnownURLs []string
}

func NewShopify() *Shopify {
	return &Shopify{URL: "www.shopify.com"}
}

func (s *Shopify) sitemapURL() string {
	return fmt.Sprintf("https://%s/sitemap.xml", s.URL)
}

func (s *Shopify) Execute() {
	c := colly.NewCollector(colly.AllowedDomains(s.URL))

	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		s.KnownURLs = append(s.KnownURLs, e.Text)
	})

	c.Visit(s.sitemapURL())

	fmt.Printf("Shopify - Known urls (Length: %v): \n", len(s.KnownURLs))
	for index, url := range s.KnownURLs {
		if index == 10 {
			break
		}
		fmt.Printf("\n %v: %s", index+1, url)
	}
}
