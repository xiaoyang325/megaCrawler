package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cprindia", "印度政策研究中心", "https://cprindia.org/")

	w.SetStartingUrls([]string{"https://cprindia.org/sitemap_index.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if ctx.Url == "https://cprindia.org/post-sitemap.xml" {
			w.Visit(element.Text, crawlers.News)
		}
		if ctx.Url == "https://cprindia.org/briefsreports-sitemap.xml" {
			w.Visit(element.Text, crawlers.Report)
		}
		if ctx.Url == "https://cprindia.org/sitemap_index.xml" {
			w.Visit(element.Text, crawlers.Index)
		}
		if ctx.Url == "https://cprindia.org/people-sitemap.xml" {
			w.Visit(element.Text, crawlers.Expert)
		}
	})

	w.OnHTML(".col-md-12 > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Title = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Name = crawlers.StandardizeSpaces(element.Text)
		}
	})

	w.OnHTML(".blog-left-text > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = element.Text
	})

	w.OnHTML(".tages-list > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML(".blog-sec > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text + "\n"
	})

	w.OnHTML(".blog-sec", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.ChildAttrs("img", "src")...)
	})

	w.OnHTML("div.blog-right-text > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".pbr-heading-sec > div > div > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.ChildText("h2")
		ctx.Authors = strings.Split(element.ChildText("h3"), "\n")
		ctx.Area = element.ChildText("p:nth-child(3)")
		ctx.PublicationTime = element.ChildText("p:nth-child(4)")
	})

	w.OnHTML(".pdf-link > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	w.OnHTML(".pbr-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".book-img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})

	w.OnHTML(".facylty-degination", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".faculty-img > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})

	w.OnHTML(".faculty-mc", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
}
