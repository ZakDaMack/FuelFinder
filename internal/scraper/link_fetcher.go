package scraper

import (
	"log/slog"
	"main/internal/slug"

	"github.com/gocolly/colly"
)

func GetTableLinks(url string) ([]string, error) {

	// make a set for unique urls
	urlSet := make(map[string]string)

	// init colly, dont limit websites as fuel data elsewhere
	c := colly.NewCollector()

	// find the data table and collect urls
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			slug := slug.Slugify(el.ChildText("td:nth-child(1)"))
			url := el.ChildText("td:nth-child(2)")
			urlSet[slug] = url
		})
	})

	// on req, ensure visiting correct page
	c.OnRequest(func(r *colly.Request) {
		slog.Debug("Visiting url", "url", r.URL)
	})

	// get any errors
	c.OnError(func(r *colly.Response, err error) {
		slog.Error("error on response", "url", r.Request.URL, "error", err)
	})

	c.Visit(url)

	// convert map to list of values
	var results []string
	for _, val := range urlSet {
		results = append(results, val)
	}

	return results, nil
}
