package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cfr", "外交关系委员会", "https://www.cfr.org/")
	w.SetStartingURLs([]string{"https://www.cfr.org/articles/sitemap.xml", "https://www.cfr.org/bio_pages/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/article/") {
			w.Visit(element.Text, crawlers.News)
		} else if strings.Contains(element.Text, "/expert/") {
			w.Visit(element.Text, crawlers.Expert)
		}
	})

	// 姓名
	w.OnHTML(".header-expert__name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})

	// 头衔
	w.OnHTML(".header-expert__dek", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 领域
	w.OnHTML(".header-expert__expertise-list > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area += element.Text + " "
	})

	// 描述,新闻正文
	w.OnHTML(".body-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Description += element.Text
		} else if ctx.PageType == crawlers.News {
			ctx.Content += element.Text
		}
	})

	// 新闻标题
	w.OnHTML(".article-header__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 作者
	w.OnHTML(".article-header__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 时间
	w.OnHTML(".article-header__date-ttr", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
}
