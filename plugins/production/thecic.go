package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("thecic", "国际理事会", "https://thecic.org/")

	w.SetStartingURLs([]string{"https://thecic.org/post-sitemap.xml"})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Authors(ctx, element)
		extractors.Titles(ctx, element)
		extractors.PublishingDate(ctx, element)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.et_pb_button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.News)
	})

	// report .content
	w.OnHTML("div.et_pb_row_1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
