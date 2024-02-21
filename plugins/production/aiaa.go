package production

import (
	"regexp"
	"strconv"
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("aiaa", "美国航空航天学会", "https://www.aiaa.org/")

	w.SetStartingURLs([]string{
		"https://www.aiaa.org/news/industry-news/1",
		"https://www.aiaa.org/news/press-releases/1",
		"https://www.aiaa.org/news/aiaa-news/1",
		"https://www.aiaa.org/publications/aerospace-america",
		"https://www.aiaa.org/events-learning/events",
		"https://www.aiaa.org/advocacy/Policy-Papers",
		"https://www.aiaa.org/get-involved/educators",
	})

	// 从翻页器获取下一页 Index 并访问
	w.OnHTML(".pagination", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		now := ctx.URL
		reg := regexp.MustCompile(`/\d+$`)
		numStr := string(reg.Find([]byte(now)))
		now = reg.ReplaceAllString(now, "")
		numStr = strings.Replace(numStr, "/", "", 1)
		num, _ := strconv.Atoi(numStr)
		num += 1
		numStr = "/" + strconv.Itoa(num)
		newURL := now + numStr
		w.Visit(newURL, crawlers.Index)
	})

	// 从 Index 访问 News
	w.OnHTML(`a.item-list__title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 从 Index 访问 News
	w.OnHTML(`a.item-list__title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(".page-title span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Publication Time
	w.OnHTML(".page-title > small:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		str := strings.Split(strings.Replace(element.Text, "Written", "", 1), "-")
		if len(str) > 1 {
			ctx.PublicationTime = strings.TrimSpace(str[1])
		} else {
			ctx.PublicationTime = strings.TrimSpace(str[0])
		}
	})

	// 获取 Content
	w.OnHTML(".group", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.ChildText("p")
	})
}
