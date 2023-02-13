package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("targetednews", "Targeted News Service", "https://targetednews.com/")

	w.SetStartingUrls([]string{
		"https://targetednews.com/daily_news.php?page=6",
		"https://targetednews.com/daily_news.php?page=10",
		"https://targetednews.com/daily_news.php?page=7",
		"https://targetednews.com/daily_news.php?page=343",
		"https://targetednews.com/daily_news.php?page=2",
		"https://targetednews.com/daily_news.php?page=4",
		"https://targetednews.com/daily_news.php?page=27",
		"https://targetednews.com/daily_news.php?page=9",
		"https://targetednews.com/daily_news.php?page=38",
		"https://targetednews.com/daily_news.php?page=19",
		"https://targetednews.com/daily_news.php?page=23",
		"https://targetednews.com/daily_news.php?page=13",
		"https://targetednews.com/daily_news.php?page=33",
		"https://targetednews.com/daily_news.php?page=18",
		"https://targetednews.com/daily_news.php?page=14",
		"https://targetednews.com/daily_news.php?page=15",
		"https://targetednews.com/daily_news.php?page=30",
		"https://targetednews.com/daily_news.php?page=12",
		"https://targetednews.com/daily_news.php?page=17",
		"https://targetednews.com/daily_news.php?page=22",
		"https://targetednews.com/newspaper_samples.php?tab=6",
		"https://targetednews.com/newspaper_samples.php?tab=1",
		"https://targetednews.com/newspaper_samples.php?tab=4",
		"https://targetednews.com/newspaper_samples.php?tab=5",
	})

	// 访问 News 从 Index 通过 SubContext
	w.OnHTML(`#content > div[id]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subContext := ctx.CreateSubContext()
		subContext.PageType = crawlers.News

		subContext.Title = element.ChildText(`div.subtitle`)
		subContext.Content = element.ChildText(`span:nth-last-child(1)`)
	})

	// 访问 News 从 Index 通过 SubContext
	w.OnHTML(`div.sample_box`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subContext := ctx.CreateSubContext()
		subContext.PageType = crawlers.News

		subContext.Title = element.ChildText(`h1`)
		subContext.PublicationTime = element.ChildText(`.dateline:nth-child(3)`)
		subContext.Content = element.ChildText(`.story_body`)
	})
}
