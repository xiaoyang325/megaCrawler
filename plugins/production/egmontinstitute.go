package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("egmontinstitute", "皇家国际关系研究所", "https://www.egmontinstitute.be/")

	w.SetStartingUrls([]string{
		"https://www.egmontinstitute.be/topics/",
		"https://www.egmontinstitute.be/publications/",
	})

	// 访问 Index 从频道入口 //
	w.OnHTML(`div.wrap main div.row div[class="row cores-container"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问下一页 Index //
	w.OnHTML(`body > div.wrap > main > div > div > div > div.row > div > div.publications > ul > li > a[class="next page-numbers"], a[title="next"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index //
	w.OnHTML(`.publications article > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title //
	w.OnHTML(`.post-publication > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`div.entry-summary > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime //
	w.OnHTML(`div.row.post-publication__header time.post-publication__date`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		time := element.Text
		time = strings.Replace(time, "(", "", 1)
		time = strings.Replace(time, ")", "", 1)
		ctx.PublicationTime = strings.TrimSpace(time)
	})

	// 获取 CategoryText //
	w.OnHTML(`div.post-publication__cat > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors //
	w.OnHTML(`div.row.post-publication__header div.post__author.post-publication__author > p > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content //
	w.OnHTML(`.post-publication__body`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags //
	w.OnHTML(`article > div.row.post-publication__header > div > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File //
	w.OnHTML(`article > div.links-container > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})
}
