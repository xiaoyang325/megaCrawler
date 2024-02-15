package production

import (
	"strings"

	"megaCrawler/extractors"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("woodson", "Woodson center", "http://woodsoncenter.org/")
	w.SetStartingURLs([]string{"http://woodsoncenter.org/post-sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.News)
	})

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

	extractorConfig.Apply(w)

	w.OnHTML(".et_pb_post_content a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "pdf") {
			ctx.Link = append(ctx.Link, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})
}
