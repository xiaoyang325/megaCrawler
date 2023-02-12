package Extractors

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
)

var rg = regexp.MustCompile(`(\r\n?|\n| ){2,}`)

func TrimText(selection *goquery.Selection) string {
	return rg.ReplaceAllString(nodeText(selection), "$1")
}

func Text(ctx *Crawler.Context, dom *colly.HTMLElement, language string) {
	node := CalculateBestNode(dom.DOM, language)
	if node == nil {
		return
	}
	node = postCleanup(node, language)
	if node == nil {
		return
	}
	if ctx.PageType == Crawler.News || ctx.PageType == Crawler.Report {
		ctx.Content = TrimText(node)
	} else if ctx.PageType == Crawler.Expert {
		ctx.Description = TrimText(node)
	}
}
