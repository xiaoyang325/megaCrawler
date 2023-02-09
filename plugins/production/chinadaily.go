package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("chinadaily", "中国日报", "https://www.chinadaily.com.cn/")
	w.SetStartingUrls([]string{
		"https://www.chinadaily.com.cn/china/governmentandpolicy",
		"https://www.chinadaily.com.cn/china/society",
		"https://www.chinadaily.com.cn/china/scitech",
		"https://www.chinadaily.com.cn/world/america/",
		"https://www.chinadaily.com.cn/world/europe",
		"https://www.chinadaily.com.cn/world/middle_east",
		"https://www.chinadaily.com.cn/world/africa",
	})

	//index
	w.OnHTML(".next > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML(".tw3_01_2_t > h4 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//标题
	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//正文
	w.OnHTML("#Content > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})
}
