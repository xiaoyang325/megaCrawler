package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cprindia", "印度政策研究中心", "https://cprindia.org/")

	w.SetStartingUrls([]string{"https://cprindia.org/sitemap_index.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if ctx.Url == "https://cprindia.org/post-sitemap.xml" {
			w.Visit(element.Text, Crawler.News)
		}
		if ctx.Url == "https://cprindia.org/briefsreports-sitemap.xml" {
			w.Visit(element.Text, Crawler.Report)
		}
		if ctx.Url == "https://cprindia.org/sitemap_index.xml" {
			w.Visit(element.Text, Crawler.Index)
		}
		if ctx.Url == "https://cprindia.org/people-sitemap.xml" {
			w.Visit(element.Text, Crawler.Expert)
		}
	})

	w.OnHTML(".col-md-12 > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = Crawler.StandardizeSpaces(element.Text)
		}
	})

	w.OnHTML(".blog-left-text > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = element.Text
	})

	w.OnHTML(".tages-list > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML(".blog-sec > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text + "\n"
	})

	w.OnHTML(".blog-sec", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = append(ctx.Image, element.ChildAttrs("img", "src")...)
	})

	w.OnHTML("div.blog-right-text > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".pbr-heading-sec > div > div > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.ChildText("h2")
		ctx.Authors = strings.Split(element.ChildText("h3"), "\n")
		ctx.Area = element.ChildText("p:nth-child(3)")
		ctx.PublicationTime = element.ChildText("p:nth-child(4)")
	})

	w.OnHTML(".pdf-link > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	w.OnHTML(".pbr-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".book-img", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})

	w.OnHTML(".facylty-degination", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".faculty-img > img", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})

	w.OnHTML(".faculty-mc", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
}
