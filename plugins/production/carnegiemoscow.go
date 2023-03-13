package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("carnegiemoscow", "卡内基莫斯科中心", "https://carnegiemoscow.org/")

	w.SetStartingURLs([]string{"https://carnegiemoscow.org/programs/74",
		"https://carnegiemoscow.org/programs/68",
		"https://carnegiemoscow.org/programs/70",
		"https://carnegiemoscow.org/programs/72",
		"https://carnegiemoscow.org/programs/1492",
		"https://carnegiemoscow.org/programs/73",
		"https://carnegiemoscow.org/programs/69",
		"https://carnegiemoscow.org/specialprojects/bookreviews?lang=en",
		"https://carnegiemoscow.org/specialprojects/russiaeudialogue/?lang=en",
		"https://carnegiemoscow.org/specialprojects/securityineurope/?lang=en",
		"https://carnegiemoscow.org/specialprojects/paxsinica/?lang=en",
		"https://carnegiemoscow.org/specialprojects/sinorussianentente/?lang=en",
		"https://carnegiemoscow.org/specialprojects/twentyfirstcenturystrategicstability/?lang=en",
		"https://carnegiemoscow.org/specialprojects/relaunchingusrussiadialogueonglobalchallenges/?lang=en",
		"https://carnegiemoscow.org/publications/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("div.content>div.center>a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("div.page-links>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从index访问新闻
	w.OnHTML("li.clearfix>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("li.clearfix>h4>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// report.title
	w.OnHTML("h1.headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	w.OnHTML(" div.col-70.tablet-zero > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	// report.author
	w.OnHTML("div.post-author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("a.em", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// report.publish time
	w.OnHTML("div.post-date>ul.clean-list>li:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML("div.headline>div>div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.category
	w.OnHTML("a.russian-text-treatment", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = element.Text
	})

	// report.description
	w.OnHTML("div.zone-title__summary", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
	// report .content
	w.OnHTML("div.article-body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	w.OnHTML("body > div.commentary-detail > div.zone-main> div.cols.no-gutters > div.zone-1.col.col-60.mobile-zero", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
