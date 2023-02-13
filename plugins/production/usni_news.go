package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("usni", "海军研究所", "https://news.usni.org/")

	w.SetStartingUrls([]string{"https://news.usni.org/",
		"https://news.usni.org/category/documents",
		"https://news.usni.org/topstories",
		"https://news.usni.org/tag/coronavirus",
		"https://news.usni.org/category/fleet-tracker"})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Image(ctx, element)
		extractors.Authors(ctx, element)
	})
	// 从翻页器获取链接并访问
	w.OnHTML("ol.wp-paginate>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	// 从index访问新闻
	w.OnHTML("div.entry-content>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// new.title
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType != crawlers.Expert {
			ctx.Title = element.Text
		}
	})
	// new.publish time
	w.OnHTML(".entry-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	// new.author
	w.OnHTML("a[rel=\"author\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = crawlers.Unique(append(ctx.Authors, element.Text))
	})
	// new.content
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	// 访问expert
	w.OnHTML("a[rel=\"author\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})
	// expert.description
	w.OnHTML("div.author-description", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	w.OnHTML(".author-description > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = strings.TrimPrefix(element.Text, "About ")
	})
	//
}
