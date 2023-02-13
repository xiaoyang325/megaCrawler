package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("bjreview", "北京周报", "https://www.bjreview.com/")
	w.SetStartingURLs([]string{"https://www.bjreview.com/China/",
		"https://www.bjreview.com/Opinion/Governance/",
		"https://www.bjreview.com/Opinion/Pacific_Dialogue/",
		"https://www.bjreview.com/Opinion/Fact_Check/",
		"https://www.bjreview.com/Opinion/Voice/",
		"https://www.bjreview.com/World/",
		"https://www.bjreview.com/Business/",
		"https://www.bjreview.com/Lifestyle/"})

	// index
	w.OnHTML("div:nth-child(1) > div > table > tbody > tr > td > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问新闻
	w.OnHTML(" li > div > table > tbody > tr > td:nth-child(2) > table > tbody > tr > td > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 标题
	w.OnHTML("td#TRS_Editor_title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 正文
	w.OnHTML(" div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
