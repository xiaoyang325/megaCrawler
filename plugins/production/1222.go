package production

import (
	"bytes"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1222", "麦克唐纳-劳里埃研究所", "https://macdonaldlaurier.ca")

	engine.SetStartingURLs([]string{"https://macdonaldlaurier.ca/sitemap_index.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		if strings.Contains(ctx.URL, ".xml") {
			dom, err := xmlquery.Parse(bytes.NewReader(bytes.TrimSpace(response.Body)))
			if err != nil {
				crawlers.Sugar.Error(err)
				return
			}
			xmlquery.FindEach(dom, "//loc", func(i int, node *xmlquery.Node) {
				var text = node.InnerText()
				switch {
				case strings.Contains(ctx.URL, "post-sitemap"):
					engine.Visit(text, crawlers.News)
				case strings.Contains(ctx.URL, "cm-expert-sitemap"):
					engine.Visit(text, crawlers.Expert)
				case strings.Contains(ctx.URL, "sitemap_index"):
					engine.Visit(text, crawlers.Index)
				}
			})
		}
	})
}
