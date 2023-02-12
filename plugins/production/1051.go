package production

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	engine := Crawler.Register("1051", "经济事务学会", "https://iea.org.uk")

	engine.SetStartingUrls([]string{"https://iea.org.uk/sitemap_index.xml"})

	extractorConfig := Extractors.Config{
		Image:       true,
		Language:    true,
		PublishDate: true,
		Tags:        true,
		Title:       true,
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "post-sitemap") {
			engine.Visit(element.Text, Crawler.Index)
		}
		if strings.Contains(ctx.Url, "post-sitemap") {
			engine.Visit(element.Text, Crawler.News)
		}
	})

	engine.OnHTML(".ph-excerpt", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		element.DOM.Find(".ph-names").Each(func(i int, selection *goquery.Selection) {
			ctx.Authors = append(ctx.Authors, selection.Text())
		})

		element.DOM.Find(".ph-author-info").Remove()

		ctx.Content = Extractors.TrimText(element.DOM)
	})
}
