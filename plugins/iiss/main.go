package iiss

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("iiss", "国际战略研究所", "https://www.iiss.org/")

	s.UrlProcessor.OnXML("//urlset/url", func(e *colly.XMLElement) {
		k, err := time.Parse("2006-01-02", e.ChildText("lastmod"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", e.ChildText("lastmod"), err.Error())
			k = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), k)
	}).SetStartingUrls([]string{"https://www.iiss.org/sitemap.xml"})

	s.UrlProcessor.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("content"))
	})

	s.UrlProcessor.OnHTML(".richtext", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("content", element.Text)
	})

	s.UrlProcessor.OnHTML(".person__name", func(element *colly.HTMLElement) {
		str := element.Request.Ctx.Get("author")
		element.Request.Ctx.Put("content", str+"\n"+element.Text)
	})
}
