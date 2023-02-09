package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"regexp"
	"strconv"
)

func init() {
	w := Crawler.Register("spri", "安全政策改革研究所", "https://www.securityreform.org/")
	w.SetStartingUrls([]string{"https://www.securityreform.org/news-and-analysis"})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		Extractors.Tags(ctx, element)
		Extractors.Titles(ctx, element)
	})

	//访问新闻
	w.OnHTML("a.summary-title-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取新闻时间
	w.OnHTML(" header > div > span.date > a > time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//获取正文
	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".like-count", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		reg, _ := regexp.Compile("\\d+")
		if val, err := strconv.Atoi(reg.FindString(element.Text)); err != nil {
			ctx.LikeCount = val
		}
	})

	w.OnHTML(".comment-count", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		reg, _ := regexp.Compile("\\d+")
		if val, err := strconv.Atoi(reg.FindString(element.Text)); err != nil {
			ctx.LikeCount = val
		}
	})
}
