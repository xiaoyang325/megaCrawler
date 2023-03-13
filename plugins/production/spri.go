package production

import (
	"regexp"
	"strconv"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("spri", "安全政策改革研究所", "https://www.securityreform.org/")
	w.SetStartingURLs([]string{"https://www.securityreform.org/news-and-analysis"})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Tags(ctx, element)
		extractors.Titles(ctx, element)
	})

	// 访问新闻
	w.OnHTML("a.summary-title-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取新闻时间
	w.OnHTML(" header > div > span.date > a > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 获取正文
	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".like-count", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		reg := regexp.MustCompile(`\d+`)
		if val, err := strconv.Atoi(reg.FindString(element.Text)); err != nil {
			ctx.LikeCount = val
		}
	})

	w.OnHTML(".comment-count", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		reg := regexp.MustCompile(`\d+`)
		if val, err := strconv.Atoi(reg.FindString(element.Text)); err != nil {
			ctx.LikeCount = val
		}
	})
}
