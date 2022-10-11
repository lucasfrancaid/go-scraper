package scraper

type Result struct {
	Data interface{}
}

type Scraper interface {
	Execute() Result
}
