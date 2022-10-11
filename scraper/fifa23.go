package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type PlayerRatings struct {
	Rank        int    `json:"rank"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
	Team        string `json:"team"`
	Position    string `json:"position"`
	Overage     int    `json:"overage"`
	Potential   int    `json:"potential"`
	URL         string `json:"url"`
}

type Fifa23 struct {
	URL     string
	Ratings []PlayerRatings
}

func NewFifa23() *Fifa23 {
	return &Fifa23{URL: "www.fifaindex.com"}
}

func (f *Fifa23) ratingsURL() string {
	return fmt.Sprintf("https://%s", f.URL)
}

func (f *Fifa23) Execute() Result {
	c := colly.NewCollector(colly.AllowedDomains(f.URL))

	// TODO: Add pagination ref: last li.page-item

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		if e.Attr("data-playerid") == "" {
			return
		}

		age, _ := strconv.Atoi(e.ChildText("td[data-title=Age]"))
		ovr, _ := strconv.Atoi(e.ChildText("td[data-title='OVR / POT'] span.badge.badge-dark.rating.r2"))
		pot, _ := strconv.Atoi(e.ChildText("td[data-title='OVR / POT'] span.badge.badge-dark.rating.r1"))
		pos := strings.Join(e.ChildTexts("td[data-title='Preferred Positions'] a span"), "|")

		p := PlayerRatings{
			Rank:        len(f.Ratings) + 1,
			Name:        e.ChildText("td[data-title=Name] a"),
			Age:         age,
			Nationality: e.ChildAttr("td[data-title=Nationality] a", "title"),
			Team:        strings.TrimSuffix(e.ChildAttr("td[data-title=Team] a", "title"), " FIFA 23"),
			Position:    pos,
			Overage:     ovr,
			Potential:   pot,
			URL:         e.ChildAttr("td[data-title=Name] a", "href"),
		}

		f.Ratings = append(f.Ratings, p)
	})

	c.Visit(f.ratingsURL())

	f.cliPrint()

	r := Result{Data: f.Ratings}
	return r
}

func (f *Fifa23) cliPrint() {
	fmt.Println("Fifa 23 - Players Rating")

	for _, p := range f.Ratings {
		fmt.Printf("\n%v° - %s, Age: %v, Pos: %s, OVR: %v, POT: %v, URL: %s",
			p.Rank, p.Name, p.Age, p.Position, p.Overage, p.Potential, f.ratingsURL()+p.URL)
	}

	players := make(map[string]interface{})
	players["ratings"] = f.Ratings

	content, err := json.Marshal(players)
	if err != nil {
		return
	}

	_ = ioutil.WriteFile("ratings.json", content, 0600)
}
