package scraper_test

import (
	"testing"

	"github.com/lucasfrancaid/go-scraper/scraper"
	"github.com/stretchr/testify/assert"
)

func TestSet_UnknownNameShouldReceiveError(t *testing.T) {
	_, err := scraper.Set("unknown")

	assert.ErrorContains(t, err, "Scraper 'unknown' not implemented")
}

func TestSet_Shopify(t *testing.T) {
	s, err := scraper.Set("shopify")

	assert.Nil(t, err)
	assert.IsType(t, &scraper.Shopify{}, s)
}

func TestSet_Fifa23(t *testing.T) {
	s, err := scraper.Set("fifa23")

	assert.Nil(t, err)
	assert.IsType(t, &scraper.Fifa23{}, s)
}
