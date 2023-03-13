package errors

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1207", "沃尔特里德国家军事医学中心", "https://walterreed.tricare.mil")

	engine.SetStartingURLs([]string{"https://walterreed.tricare.mil/SiteMap.aspx"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/News-Gallery/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})
}
