package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("dfa_gov_ph", "外交部外交服务研究所", "https://dfa.gov.ph/")

	w.SetStartingURLs([]string{
		"https://dfa.gov.ph/dfa-news/dfa-releasesupdate",
		"https://dfa.gov.ph/dfa-news/statements-and-advisoriesupdate",
		"https://dfa.gov.ph/dfa-news/news-from-our-foreign-service-postsupdate",
		"https://dfa.gov.ph/dfa-news/events/kalayaan-2022",
		"https://dfa.gov.ph/dfa-news/events/araw-ng-kagitingan-2022-1",
		"https://dfa.gov.ph/dfa-news/events/rizal-day-2020",
		"https://dfa.gov.ph/dfa-news/events/year-end-holiday-2019",
		"https://dfa.gov.ph/dfa-news/events/kalayaan-2019",
		"https://dfa.gov.ph/dfa-news/events/christmas-2018",
		"https://dfa.gov.ph/dfa-news/events/kalayaan-2018",
		"https://dfa.gov.ph/dfa-news/events/women-s-month",
		"https://dfa.gov.ph/dfa-news/events/asean-50-celebration",
		"https://dfa.gov.ph/dfa-news/events/kalayaan-2017",
		"https://dfa.gov.ph/dfa-news/events/54th-asean-ministerial-meeting",
		"https://dfa.gov.ph/dfa-news/events/christmas-2017",
		"https://dfa.gov.ph/dfa-news/events/rizal-day-2017",
	})

	// 访问下一页 Index
	w.OnHTML(`.pagination-next > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`tbody > tr > td > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(`.entry-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`#content > div.post-box.clearfix > div.item-page > div.page-header > h2`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`#content > div.post-box.clearfix > div.item-page > div[itemprop="articleBody"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
