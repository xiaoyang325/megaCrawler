package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cepa", "欧洲政策分析中心", "https://cepa.org/")
	w.SetStartingUrls([]string{"https://cepa.org/about-cepa/team/experts/", "https://cepa.org/insights-analysis/"})

	// 访问专家
	w.OnHTML(".btn.btn-primary.text-cta", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 获取专家姓名
	w.OnHTML("div.flex-col.gap-\\[14px\\].hidden.md\\:flex > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})

	// 专家头衔
	w.OnHTML("div.flex-col.gap-\\[14px\\].hidden.md\\:flex > h5", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 专家领域
	w.OnHTML("div.flex.flex-col.gap-\\[14px\\].flex-1 > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area += element.Text + " "
	})

	// 专家描述
	w.OnHTML(".post-content.\\[\\&\\>\\*\\]\\:leading-\\[26px\\].transition-all.ease-in.duration-1000 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	// 访问新闻
	w.OnHTML(".group.cursor-pointer", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 新闻标题
	w.OnHTML(".text-h1.text-dark-blue", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Title = element.Text
		}
	})

	// 新闻标签
	w.OnHTML(" div.flex.gap-\\[14px\\].flex-wrap > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 新闻时间
	w.OnHTML("div.flex.flex-col.gap-\\[9px\\].text-blue.font-bold > div:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 作者
	w.OnHTML("div.\\[\\&\\>a\\:hover\\]\\:text-red", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 正文
	w.OnHTML(".post-content.post-container.mt-\\[110px\\]>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
