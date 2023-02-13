package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("tibetaction", "美国西藏行动委员会", "https://tibetaction.net/")
	w.SetStartingURLs([]string{"https://tibetaction.net/wp-sitemap-posts-post-1.xml", "https://tibetaction.net/wp-sitemap-posts-profile-1.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(ctx.URL, "wp-sitemap-posts-post") {
			w.Visit(element.Text, crawlers.News)
		} else if strings.Contains(ctx.URL, "wp-sitemap-posts-profile") {
			w.Visit(element.Text, crawlers.Expert)
		}
	})

	w.OnHTML(".published", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".entry-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Title = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".cmsmasters_profile_subtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".cmsmasters_img img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	w.OnHTML(".cmsmasters_text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
}
