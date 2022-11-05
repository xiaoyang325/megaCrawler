package carnegieendowment

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("carnegieendowment", "卡内基国际和平基金会",
									  "https://carnegieendowment.org/")

	w.SetStartingUrls([]string{
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
	w.OnHTML("div[class=\"section research\"]>.foreground>.center>a[class=\"button teal\"]",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			url := "https://carnegieendowment.org/" + element.Attr("href")
			w.Visit(url, megaCrawler.Index)
		})

	// 从翻页器获取下一页 Index 并访问
	w.OnHTML("div[class=\"center section\"]>div>a[class=\"page-links__next tag uppercase\"]",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			url := "https://carnegieendowment.org/" + element.Attr("href")
			w.Visit(url, megaCrawler.Index)
		})


	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	// 从 Index 访问 Report
	w.OnHTML(".clearfix>.no-margin>a",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			w.Visit(element.Attr("href"), megaCrawler.Report)
		})

	// 添加 Title 到 ctx
	w.OnHTML(".headline",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 添加 Author 到 ctx
	w.OnHTML("div[class=\"post-author col col-75\"]>a", 
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 添加 Author 到 ctx
	w.OnHTML("div[class=\"post-author col col-75\"]", 
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 添加 Content 到 ctx
	w.OnHTML("div.article-body", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 添加 Tag 到 ctx
	w.OnHTML(".show-tag", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			ctx.Tags = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

}
