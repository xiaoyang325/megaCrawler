package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1051", "经济事务学会", "https://iea.org.uk")

	engine.SetStartingURLs([]string{"https://iea.org.uk/sitemap_index.xml"})

	extractorConfig := extractors.Config{
		Image:       true,
		Language:    true,
		PublishDate: true,
		Tags:        true,
		Title:       true,
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		}
		if strings.Contains(ctx.URL, "post-sitemap") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".ph-excerpt", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".ph-names").Each(func(i int, selection *goquery.Selection) {
			ctx.Authors = append(ctx.Authors, selection.Text())
		})

		element.DOM.Find(".ph-author-info").Remove()

		ctx.Content = extractors.TrimText(element.DOM)
	})
}
