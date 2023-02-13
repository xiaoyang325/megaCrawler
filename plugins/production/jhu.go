package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("jhu", "Stavros Niarchos Foundation Agora Institute", "https://snfagora.jhu.edu/")
	w.SetStartingURLs([]string{"https://snfagora.jhu.edu/wp-sitemap-posts-person-1.xml", "https://snfagora.jhu.edu/wp-sitemap-posts-news-1.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, crawlers.News)
		}
		if strings.Contains(element.Text, "/person/") {
			w.Visit(element.Text, crawlers.Expert)
		}
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Title = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".article__sub-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.SubTitle = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Title = element.Text
		}
	})

	w.OnHTML(".article__content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Content = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Description = element.Text
		}
	})

	w.OnHTML(".article__date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
}
