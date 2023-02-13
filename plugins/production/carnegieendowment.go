package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("carnegieendowment", "卡内基国际和平基金会",
		"https://carnegieendowment.org/")

	w.SetStartingURLs([]string{
		"https://carnegieendowment.org/programs/africa",
		"https://carnegieendowment.org/programs/americanstatecraft/",
		"https://carnegieendowment.org/programs/asia/",
		"https://carnegieendowment.org/programs/democracy/",
		"https://carnegieendowment.org/programs/europe/",
		"https://carnegieendowment.org/programs/globalorder/",
		"https://carnegieendowment.org/programs/middleeast/",
		"https://carnegieendowment.org/programs/npp/",
		"https://carnegieendowment.org/programs/russia/",
		"https://carnegieendowment.org/programs/southasia/",
		"https://carnegieendowment.org/programs/climate/",
		"https://carnegieendowment.org/programs/technology/",
	})

	// 从子频道入口访问 All Program Publications
	w.OnHTML("div[class=\"section research\"]>.foreground>.center>a[class=\"button teal\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从翻页器获取下一页 Index 并访问
	w.OnHTML("div[class=\"center section\"]>div>a[class=\"page-links__next tag uppercase\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从 Index 访问 Report
	w.OnHTML(".clearfix>.no-margin>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 添加 Title 到 ctx
	w.OnHTML(".headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 添加 Author 到 ctx （第二种情况）
	w.OnHTML("div[class=\"post-author col col-75\"]>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 添加 Author 到 ctx （第一种情况）
	w.OnHTML("div[class=\"post-author col col-75\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 添加 Content 到 ctx
	w.OnHTML("div.article-body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 添加 Tag 到 ctx
	w.OnHTML(".show-tag", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
