package production

import (
	"strconv"
	"strings"
	"time"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("news_fit", "佛罗里达理工学院", "https://news.fit.edu/")

	w.SetStartingURLs([]string{
		"https://news.fit.edu/sitemap_index.xml",
	})

	// 访问 sitemap
	w.OnXML(`//loc`, func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, ".xml") {
			w.Visit(element.Text, crawlers.Index)
		} else {
			w.Visit(element.Text, crawlers.News)
		}
	})

	// 获取 Title
	w.OnHTML(`[class="post-title entry-title"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`.entry-sub-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.entry-header [class="date meta-item fa-before"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		split := strings.Split(strings.TrimSpace(element.Text), " ")
		if len(split) < 2 {
			return
		}
		now := time.Now()
		unit := split[1]
		number, err := strconv.Atoi(split[0])
		if err != nil {
			return
		}
		if unit == "day" || unit == "days" {
			now = now.AddDate(0, 0, -number)
		}
		if unit == "week" || unit == "weeks" {
			now = now.AddDate(0, 0, -number*14)
		}
		if unit == "month" || unit == "months" {
			now = now.AddDate(0, -number, 0)
		}
		if unit == "year" || unit == "years" {
			now = now.AddDate(-number, 0, 0)
		}
		ctx.PublicationTime = now.Format(time.RFC3339)
	})

	// 获取 Authors
	w.OnHTML(`.meta-author .author-name`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`[class="entry-content entry clearfix"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})
}
