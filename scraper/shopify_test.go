package scraper_test

import (
	"testing"

	"github.com/lucasfrancaid/go-scraper/scraper"
	"github.com/stretchr/testify/assert"
)

func TestShopify_ImplementsScraperInterface(t *testing.T) {
	assert.Implements(t, (*scraper.Scraper)(nil), new(scraper.Shopify))
}

func TestShopify_NewShopify(t *testing.T) {
	s := scraper.NewShopify()

	assert.IsType(t, &scraper.Shopify{}, s)
	assert.NotNil(t, s.URL)
	assert.Nil(t, s.KnownURLs)
}
