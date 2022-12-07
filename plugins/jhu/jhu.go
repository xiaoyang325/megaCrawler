package jhu

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("jhu", "Stavros Niarchos Foundation Agora Institute", "https://snfagora.jhu.edu/")
	w.SetStartingUrls([]string{"https://snfagora.jhu.edu/wp-sitemap-posts-person-1.xml", "https://snfagora.jhu.edu/wp-sitemap-posts-news-1.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, Crawler.News)
		}
		if strings.Contains(element.Text, "/person/") {
			w.Visit(element.Text, Crawler.Expert)
		}
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".article__sub-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.SubTitle = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Title = element.Text
		}
	})

	w.OnHTML(".article__content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Content = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		}
	})

	w.OnHTML(".article__date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
