package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("tibetaction", "美国西藏行动委员会", "https://tibetaction.net/")
	w.SetStartingUrls([]string{"https://tibetaction.net/wp-sitemap-posts-post-1.xml", "https://tibetaction.net/wp-sitemap-posts-profile-1.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(ctx.Url, "wp-sitemap-posts-post") {
			w.Visit(element.Text, Crawler.News)
		} else if strings.Contains(ctx.Url, "wp-sitemap-posts-profile") {
			w.Visit(element.Text, Crawler.Expert)
		}
	})

	w.OnHTML(".published", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".cmsmasters_profile_subtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".cmsmasters_img img", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	w.OnHTML(".cmsmasters_text", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
}
