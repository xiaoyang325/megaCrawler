package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	w := Crawler.Register("thecic", "国际理事会", "https://thecic.org/")

	w.SetStartingUrls([]string{"https://thecic.org/post-sitemap.xml"})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		Extractors.Authors(ctx, element)
		Extractors.Titles(ctx, element)
		Extractors.PublishingDate(ctx, element)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.et_pb_button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		w.Visit(element.Text, Crawler.News)
	})

	// report .content
	w.OnHTML("div.et_pb_row_1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
