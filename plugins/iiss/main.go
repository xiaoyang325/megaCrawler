package iiss

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("iiss", "https://www.iiss.org/")
	s.UrlProcessor.OnXML("//urlset/url", func(e *colly.XMLElement) {
		k, err := time.Parse("2006-01-02", e.ChildText("lastmod"))
		if err != nil {
			if err != nil {
				megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", e.ChildText("lastmod"), err.Error())
				k = time.Unix(0, 0)
			}
		}
		s.AddUrl(e.ChildText("loc"), k)
	}).SetStartingUrls([]string{"/sitemap.xml/"})

	s.UrlProcessor.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("content"))
	})

	s.UrlProcessor.OnHTML(".richtext > span > p > span", func(element *colly.HTMLElement) {
		script := element.ChildText("script")
		if script != "" {
			return
		}
		str := element.Request.Ctx.Get("content")
		element.Request.Ctx.Put("content", str+"\n"+element.Text)
	})
}
