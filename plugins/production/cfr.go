package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cfr", "外交关系委员会", "https://www.cfr.org/")
	w.SetStartingUrls([]string{"https://www.cfr.org/articles/sitemap.xml", "https://www.cfr.org/bio_pages/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "/article/") {
			w.Visit(element.Text, Crawler.News)
		} else if strings.Contains(element.Text, "/expert/") {
			w.Visit(element.Text, Crawler.Expert)
		}
	})

	//姓名
	w.OnHTML(".header-expert__name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//头衔
	w.OnHTML(".header-expert__dek", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//领域
	w.OnHTML(".header-expert__expertise-list > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area += element.Text + " "
	})

	//描述,新闻正文
	w.OnHTML(".body-content > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Description += element.Text
		} else if ctx.PageType == Crawler.News {
			ctx.Content += element.Text
		}
	})

	//新闻标题
	w.OnHTML(".article-header__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//作者
	w.OnHTML(".article-header__link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//时间
	w.OnHTML(".article-header__date-ttr", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
