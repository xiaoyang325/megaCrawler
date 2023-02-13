package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("eliamep", "希腊欧洲与外交政策基金会", "https://www.eliamep.gr/en/")
	w.SetStartingURLs([]string{"https://www.eliamep.gr/en/publications/",
		"https://www.eliamep.gr/en/experts/"})

	// index
	w.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问专家
	w.OnHTML("a.name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 访问报告
	w.OnHTML("a.title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 报告标题,专家姓名
	w.OnHTML("h1.postTitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Report {
			ctx.Title = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	// 专家头衔
	w.OnHTML("h2.postSubTitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 报告时间
	w.OnHTML("div.l2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 报告正文,专家描述
	w.OnHTML(".articleBody", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Report {
			ctx.Content = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Description = element.Text
		}
	})
}
