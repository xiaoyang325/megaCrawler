package main

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func main() {
	w := megaCrawler.Register("cato", "卡托研究所", "https://www.cato.org")

	w.OnHTML("a.btn", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.PageType = megaCrawler.Report
		}

	})

	megaCrawler.Start()
}
