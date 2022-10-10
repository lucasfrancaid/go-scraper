package scraper_test

import (
	"testing"

	"github.com/lucasfrancaid/go-scraper/scraper"
	"github.com/stretchr/testify/assert"
)

func TestFifa23_ImplementsScraperInterface(t *testing.T) {
	assert.Implements(t, (*scraper.Scraper)(nil), new(scraper.Fifa23))
}

func TestFifa23_NewFifa23(t *testing.T) {
	s := scraper.NewFifa23()

	assert.IsType(t, &scraper.Fifa23{}, s)
	assert.NotNil(t, s.URL)
	assert.Nil(t, s.Ratings)
}
