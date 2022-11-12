package ips

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("ips", "伊斯兰堡政策研究所", "https://www.ips.org.pk/")
	w.SetStartingUrls([]string{"https://www.ips.org.pk/category/ips-events/",
		"https://www.ips.org.pk/ips-lead/the-living-scripts/",
		"https://www.ips.org.pk/category/research/pakistan-affairs/",
		"https://www.ips.org.pk/category/research-themes/international-relations/",
		"https://www.ips.org.pk/category/research-themes/faith-and-society/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("div.entry-content > table > tbody > tr > td > h3 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})
	//访问index
	w.OnHTML(" div > div.pagination-wrap > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	//从index中访问文章
	w.OnHTML(" div.post-content > h2 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})
	w.OnHTML(" div > div> div > h4 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//获取文章标题
	w.OnHTML("#content > article > div.post-content > h2", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	//获取文章标签
	w.OnHTML(" div.post-content > div.post-meta > span.meta-cats > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	//获取文章图片
	w.OnHTML(" div.owl-stage-outer.owl-height > div > div > div > div > img", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	//获取文章正文
	w.OnHTML("article > div.post-content > div.entry-content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML("meta[property=\"article:published_time\"]", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
