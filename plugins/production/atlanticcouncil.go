package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"regexp"
	"strings"
)

func init() {
	w := Crawler.Register("atlanticcouncil", "大西洋理事會", "https://www.atlanticcouncil.org/")

	w.SetStartingUrls([]string{"https://www.atlanticcouncil.org/sitemap_index.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		reg := regexp.MustCompile(`([a-zA-Z_-]+)\d*.xml`)
		switch reg.FindStringSubmatch(ctx.Url)[1] {
		case "sitemap_index":
			w.Visit(element.Text, Crawler.Index)
		case "post-sitemap":
			w.Visit(element.Text, Crawler.News)
		case "expert-sitemap":
			w.Visit(element.Text, Crawler.Expert)
		}
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		Extractors.Image(ctx, element)
	})

	w.OnHTML(".ac-single-post--content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if element.ChildText(".ac-single-post--marquee") != "" {
			ctx.Content = strings.Join(element.ChildTexts("p"), "\n")
			return
		}
		ctx.Content = Crawler.HTML2Text(strings.TrimSpace(element.Text))
	})

	w.OnHTML(".gta-post-site-banner--title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("*[class$=’heading--date’]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".gta-post-site-banner--tax--cats", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	w.OnHTML(".gta-post-embed--tax--expert", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".ac-single-post--content a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	w.OnHTML(".gta-expert-site-banner--title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	w.OnHTML(".gta-expert-site-banner--positions > li", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(ctx.Title + "\n" + element.Text)
	})

	w.OnHTML(".ac-single-expert--meta:nth-child(1)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = strings.TrimSpace(ctx.Area + "\n" + element.Text)
	})

	w.OnHTML(".ac-single-expert--meta:nth-child(2)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Location = strings.TrimSpace(ctx.Location + "\n" + element.Text)
	})

	w.OnHTML(".ac-single-expert--content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
}
