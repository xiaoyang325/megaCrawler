package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1267", "彼特·多尔蒂感染与免疫研究所", "https://www.doherty.edu.au")

	engine.SetStartingURLs([]string{"https://www.doherty.edu.au/sitemap.xml"})

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
		if strings.Contains(element.Text, "/our-work/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})
}
