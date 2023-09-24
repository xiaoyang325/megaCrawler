package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("nr", "Navy Recognition", "https://navyrecognition.com/")
	engine.SetStartingURLs([]string{"https://navyrecognition.com/index.php/naval-news.html"})

	extractorConfig := extractors.Config{
		Author:      true,
		Image:       true,
		Language:    true,
		PublishDate: true,
		Tags:        true,
		Text:        true,
		Title:       true,
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".readmore > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
}
