package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("sipri", "斯德哥尔摩国际和平研究所", "https://sipri.org/")
	w.SetStartingUrls([]string{"https://sipri.org/news/past",
		"https://sipri.org/publications/search",
		"https://sipri.org/media/issue_experts"})

	//index
	w.OnHTML(".pager__item--next > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问专家
	w.OnHTML("h2.field-content > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//访问报告
	w.OnHTML(".views-field-title > em > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//访问新闻
	w.OnHTML(".field-content > h3 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//文章标题,专家姓名
	w.OnHTML(".block-core > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News || ctx.PageType == Crawler.Report {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	//报告正文
	w.OnHTML(".node.node--type-publication.node--view-mode-full.ds-1col", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//报告pdf
	w.OnHTML(" div.field-pdf-full-publication.field--label-hidden > div > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//新闻正文
	w.OnHTML("article > div > div.body.field--label-hidden > div > div ", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//新闻时间
	w.OnHTML(" div > div > time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//作者
	w.OnHTML(" div.views-field.views-field-combinedauthors > span > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//专家领域
	w.OnHTML(" div.field-subject-expertise.clearfix.field--label-inline > div.field-items > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//专家介绍
	w.OnHTML("#sipri-2016-content > div > div.body.field--label-hidden > div > div ", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
}
