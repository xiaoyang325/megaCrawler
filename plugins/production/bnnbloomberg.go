package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("bnnbloomberg", "bnnbloomberg", "https://www.bnnbloomberg.ca/")

	w.SetStartingURLs([]string{
		"https://www.bnnbloomberg.ca/", "https://www.bnnbloomberg.ca/investing", "https://www.bnnbloomberg.ca/etfs", "https://www.bnnbloomberg.ca/personal-finance", "https://www.bnnbloomberg.ca/personal-finance/video", "https://www.bnnbloomberg.ca/company-news", "https://www.bnnbloomberg.ca/company-news/video", "https://www.bnnbloomberg.ca/commodities", "https://www.bnnbloomberg.ca/economics", "https://www.bnnbloomberg.ca/politics", "https://www.bnnbloomberg.ca/technology", "https://www.bnnbloomberg.ca/bloomberg-news-wire", "https://www.bnnbloomberg.ca/opinion", "https://www.bnnbloomberg.ca/executive",
	})

	w.OnHTML(".sources>p>.author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".article-text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a#loadmoreLatestNewsnontabbed", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML(".load-more", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML(".normal-link>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".scroll-container-frame>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".top-story-headline>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".headline-super>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".related>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".article-content>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".media-content>.headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

}
