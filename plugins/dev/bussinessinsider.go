package dev

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("businessinsider", "businessinsider", "https://www.businessinsider.in/")

	w.SetStartingURLs([]string{
		"https://www.businessinsider.in/", "https://www.businessinsider.in/business", "https://www.businessinsider.in/tech", "https://www.businessinsider.in/stock-market",
	})

	w.OnHTML("#Content > div.articlepage.clearfix > div.wrapper.clearfix.article-content-wrapper > div.box-lhs-width.float-left > div > div.article_content.clearfix > div > article > div.mobile_padding > div.social_byline_pnl.clearfix > div.ByLine > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("#Content > div.articlepage > div.wrapper.article-content-wrapper > div > div > div.article_content > div > article > div.ArtInnerCont > div.article_content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a.tout-title-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".story-headline>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("a.list-title-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("#Content > div.articlepage.clearfix > div.wrapper.clearfix.article-content-wrapper > div.box-lhs-width.float-left > div > div.article_content.clearfix > div > article > div.mobile_padding > div.social_byline_pnl.clearfix > div.ByLine > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML("div.wrapper.clearfix.article-content-wrapper > div.box-lhs-width.float-left > div > div.article_content.clearfix > div > article > div.mobile_padding > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})

}
