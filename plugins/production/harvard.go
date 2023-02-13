package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("harvard", "哈佛大学政治研究所", "https://iop.harvard.edu/")

	w.SetStartingUrls([]string{"https://iop.harvard.edu/fall-2022-harvard-youth-poll",
		"https://iop.harvard.edu/conferences",
		"https://iop.harvard.edu/youth-poll/spring-2022-harvard-youth-poll",
		"https://iop.harvard.edu/youth-poll/fall-2021-harvard-youth-poll",
		"https://iop.harvard.edu/youth-poll/spring-2021-harvard-youth-poll",
		"https://iop.harvard.edu/youth-poll/past"})

	// 从index访问新闻
	w.OnHTML("div.field-item>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// report .content
	w.OnHTML("#block-system-main > div > div.paragraphs-items.paragraphs-items-field-s-paragraph.paragraphs-items-field-s-paragraph-full.paragraphs-items-full > div > div > div:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
