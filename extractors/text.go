package extractors

import (
	"regexp"

	"megaCrawler/crawlers"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

var rg = regexp.MustCompile(`(\r\n?|\n| ){2,}`)

func TrimText(selection *goquery.Selection) string {
	return rg.ReplaceAllString(nodeText(selection), "$1")
}

func Text(ctx *crawlers.Context, dom *colly.HTMLElement, language string) {
	node := CalculateBestNode(dom.DOM, language)
	if node == nil {
		return
	}
	node = postCleanup(node, language)
	if node == nil {
		return
	}
	if ctx.PageType == crawlers.News || ctx.PageType == crawlers.Report {
		ctx.Content = TrimText(node)
	} else if ctx.PageType == crawlers.Expert {
		ctx.Description = TrimText(node)
	}
}
