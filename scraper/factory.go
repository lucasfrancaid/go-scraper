package scraper

import (
	"errors"
	"fmt"
)

func Set(name string) (Scraper, error) {
	switch name {
	case "shopify":
		return NewShopify(), nil
	case "fifa23":
		return NewFifa23(), nil
	default:
		err := fmt.Sprintf("Scraper '%v' not implemented", name)
		return nil, errors.New(err)
	}
}
