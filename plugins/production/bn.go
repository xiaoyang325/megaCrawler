package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("bn", "Benar News", "https://www.benarnews.org/")
	engine.SetStartingURLs([]string{"https://www.benarnews.org/english/news/story_archive"})

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

	engine.OnHTML("#storycontent > div:nth-child(3) > span > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("#storycontent > nav > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".archive > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
}
