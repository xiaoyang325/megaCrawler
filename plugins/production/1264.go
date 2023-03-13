package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1264", "澳大利亚战略政策研究所", "https://www.aspi.org.au")

	engine.SetStartingURLs([]string{"https://www.aspi.org.au/sitemap.xml"})

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
		switch {
		case strings.Contains(element.Text, ".xml"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "/bio/"):
			engine.Visit(element.Text, crawlers.Expert)
		case strings.Contains(element.Text, "/report/"):
			engine.Visit(element.Text, crawlers.Report)
		}
	})
}
