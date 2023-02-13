package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("siia", "国际事务研究所", "http://www.siiaonline.org/")
	w.SetStartingUrls([]string{"http://www.siiaonline.org/",
		"http://www.siiaonline.org/commentaries/",
		"http://www.siiaonline.org/insights/",
		"http://www.siiaonline.org/reports-index/"})

	// 从翻页器获取链接并访问
	w.OnHTML("a.inactive", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 新闻索引页
	w.OnHTML("a.tag-cloud-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从index访问新闻
	w.OnHTML(".post_more > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	w.OnHTML("div.post_text > div > h5 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 访问报告
	w.OnHTML("div.vc_btn3-container.vc_btn3-left > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 新闻,报告标题
	w.OnHTML(" div > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News || ctx.PageType == crawlers.Report {
			ctx.Title = element.Text
		}
	})

	// 新闻，报告图片
	w.OnHTML(" .post_image > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	// 新闻，报告标签
	w.OnHTML(".single_tags.clearfix > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 新闻，报告正文
	w.OnHTML("div.post_text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News || ctx.PageType == crawlers.Report {
			ctx.Content = element.Text
		}
	})

	// 报告pdf
	w.OnHTML(" div > p> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}
