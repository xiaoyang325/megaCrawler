package iiss

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("iiss", "International Institute for Strategic Studies",
			"https://www.iiss.org/")
	
	w.SetStartingUrls([]string{
		"https://www.iiss.org/sitemap.xml",
	})

	// 访问文章从 sitemap
	w.OnXML(`//loc`,
		func(element *colly.XMLElement, ctx *Crawler.Context) {
			if (strings.Contains(element.Text, "/blogs/")) {
				w.Visit(element.Text, Crawler.Report)
			} else if (strings.Contains(element.Text, "/press/")) {
				w.Visit(element.Text, Crawler.News)
			} else if (strings.Contains(element.Text, "/publications/")) {
				w.Visit(element.Text, Crawler.Report)
			} else if (strings.Contains(element.Text, "/events/")) {
				w.Visit(element.Text, Crawler.Report)
			} else if (strings.Contains(element.Text, "/people/")) {
				w.Visit(element.Text, Crawler.Expert)
			}
		})

	// 获取 Title 或 Name
	w.OnHTML(`div[class="col col--span6 col--span6_large col--push_1_medium container--main"] div > h1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if ctx.PageType != Crawler.Expert {
				ctx.Title = strings.TrimSpace(element.Text)
			} else {
				ctx.Name = strings.TrimSpace(element.Text)
			}
		})

	// 获取 Description 或 Expert's Title
	w.OnHTML(`div[class="col col--span6 col--span6_large col--push_1_medium container--main"] div >div.intro`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if ctx.PageType != Crawler.Expert {
				ctx.Description = strings.TrimSpace(element.Text)
			} else {
				ctx.Title = strings.TrimSpace(element.Text)
			}
		})

	// 获取 PublicationTime
	w.OnHTML(`div.article__title > p[class="label label--date"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`dl.data--highlight > dd:nth-child(3) > span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`p[class="label label--small"] > strong > span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText
	w.OnHTML(`div.article__title > a[class="label label--link"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`div.col--inner div.person p[class="h5 person__name"] span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Content 或是 Expert's Description
	w.OnHTML(`div[class="richtext component"]:nth-child(1)`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if ctx.PageType != Crawler.Expert {
				ctx.Content = strings.TrimSpace(element.ChildText("p h1 h2 h3"))
			} else {
				ctx.Description = strings.TrimSpace(element.ChildText("p h1 h2 h3"))
			}
		})

	// 获取 File
	w.OnHTML(`div[class="linklist component"] a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			file_url := "https://www.iiss.org" + element.Attr("href")
			ctx.File = append(ctx.File, file_url)
		})
}
