package Extractors

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
)

func Text(ctx *Crawler.Context, dom *colly.HTMLElement, language string) {
	node := CalculateBestNode(dom.DOM, language)
	node = postCleanup(node, language)
	if node == nil {
		return
	}
	rg := regexp.MustCompile(`(\r\n?|\n){2,}`)
	k := nodeText(node)
	println(rg.ReplaceAllString(k, "$1"))
}
