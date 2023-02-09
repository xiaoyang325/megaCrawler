package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("eliamep", "希腊欧洲与外交政策基金会", "https://www.eliamep.gr/en/")
	w.SetStartingUrls([]string{"https://www.eliamep.gr/en/publications/",
		"https://www.eliamep.gr/en/experts/"})

	//index
	w.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问专家
	w.OnHTML("a.name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//访问报告
	w.OnHTML("a.title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//报告标题,专家姓名
	w.OnHTML("h1.postTitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	//专家头衔
	w.OnHTML("h2.postSubTitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//报告时间
	w.OnHTML("div.l2", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//报告正文,专家描述
	w.OnHTML(".articleBody", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Content = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		}
	})
}
