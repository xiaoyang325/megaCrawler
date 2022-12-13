package aiaa

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	w := Crawler.Register("aiaa", "美国航空航天学会", "https://www.aiaa.org/")

	w.SetStartingUrls([]string{
		"https://www.aiaa.org/news/industry-news/1",
		"https://www.aiaa.org/news/press-releases/1",
		"https://www.aiaa.org/news/aiaa-news/1",
		"https://www.aiaa.org/publications/aerospace-america",
		"https://www.aiaa.org/events-learning/events",
		"https://www.aiaa.org/advocacy/Policy-Papers",
		"https://www.aiaa.org/get-involved/educators",
	})

	// 从翻页器获取下一页 Index 并访问
	w.OnHTML(".pagination", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		now := ctx.Url
		reg, _ := regexp.Compile(`/\d+$`)
		num_str := string(reg.Find([]byte(now)))
		now = reg.ReplaceAllString(now, "")
		num_str = strings.Replace(num_str, "/", "", 1)
		num, _ := strconv.Atoi(num_str)
		num += 1
		num_str = "/" + strconv.Itoa(num)
		new_url := now + num_str
		w.Visit(new_url, Crawler.Index)
	})

	// 从 Index 访问 News
	w.OnHTML(`a.item-list__title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// 从 Index 访问 News
	w.OnHTML(`a.item-list__title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// 获取 Title
	w.OnHTML(".page-title span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Publication Time
	w.OnHTML(".page-title > small", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		str := strings.Replace(element.Text, "Written", "", 1)
		ctx.PublicationTime = strings.TrimSpace(str)
	})

	// 获取 Content
	w.OnHTML(".group", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.ChildText("p")
	})
}
