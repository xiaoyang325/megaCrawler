package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("todayonline", "todayonline", "https://www.todayonline.com/")

	w.SetStartingURLs([]string{
		"https://www.todayonline.com/", "https://www.todayonline.com/singapore", "https://www.todayonline.com/world", "https://www.todayonline.com/big-read", "https://www.todayonline.com/adulting-101", "https://www.todayonline.com/gen-y-speaks", "https://www.todayonline.com/gen-z-speaks", "https://www.todayonline.com/voices", "https://www.todayonline.com/commentary", "https://www.todayonline.com/8days", "https://www.todayonline.com/health", "https://www.todayonline.com/brand-spotlight",
	})

	w.OnHTML("#block-mc-todayonline-theme-mainpagecontent > article:nth-child(1) > div.content > div:nth-child(3) > div.layout__region.layout__region--second > section.block.block-layout-builder.block-field-blocknodearticlefield-author.clearfix > div > div > div > div.author-card__content > div > div > h6 > a > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML(".paragraph", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML(".list-object__heading", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML(".h4__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".list-object__heading-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("#block-mc-todayonline-theme-mainpagecontent > article:nth-child(1) > div.content > div:nth-child(3) > div.layout__region.layout__region--second > section.block.block-mc-content-share-bookmark.block-content-share-bookmark.clearfix > div.article-date.article-date-- > div:nth-child(1) > font > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML(".h6--author-name>.h6__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})
	w.OnHTML(".h1--page-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
	w.OnHTML(".h1--author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name += element.Text
	})
}
