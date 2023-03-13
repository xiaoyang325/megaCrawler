package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("isgi", "国际法律研究所", "https://www.isgi.cnr.it/")

	w.SetStartingURLs([]string{"https://www.isgi.cnr.it/altri-eventi/",
		"https://www.isgi.cnr.it/progetti-conclusi-2/",
		"https://www.isgi.cnr.it/pubblicazioni/italian-reports-on-international-humanitarian-law/",
		"https://www.isgi.cnr.it/pubblicazioni/la-prassi-italiana-di-diritto-internazionale/",
		"https://www.isgi.cnr.it/pubblicazioni/marsafenet-open-access-publications/",
		"https://www.isgi.cnr.it/pubblicazioni/pubblicazioni-daic/",
		"https://www.isgi.cnr.it/pubblicazioni/altre-pubblicazioni/",
		"https://www.isgi.cnr.it/pubblicazioni/altri-volumi/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("div.post-content>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})
	w.OnHTML("div.post-content>ul>li>strong>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})
	w.OnHTML("div.post-content>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从index访问新闻
	w.OnHTML("div.post-content>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// report.title
	w.OnHTML("h1.post-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	// report.publish time
	w.OnHTML("span.published", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("div.author>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report .content
	w.OnHTML("div.post-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
