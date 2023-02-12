package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	engine := Crawler.Register("1075", "斯德哥尔摩国际和平研究所", "https://www.sipri.org")

	engine.SetStartingUrls([]string{"https://www.sipri.org/sitemap.xml"})

	extractorConfig := Extractors.Config{
		Author:      true,
		Image:       true,
		Language:    true,
		PublishDate: true,
		Tags:        true,
		Text:        true,
		Title:       true,
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "sitemap.xml") {
			engine.Visit(element.Text, Crawler.Index)
		}
		if strings.Contains(element.Text, "/about/bios/") {
			engine.Visit(element.Text, Crawler.Expert)
		}
		if strings.Contains(element.Text, "/node/") || strings.Contains(element.Text, "/news/") {
			engine.Visit(element.Text, Crawler.News)
		}
	})
}
