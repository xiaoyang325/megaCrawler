package extractors

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/gocolly/colly/v2"
)

type selectorContentPair struct {
	selector string
	content  string
}

// GetMetaContent Extract a given meta content form document.
//
//	Example metaNames:
//	   "meta[name=description]"
//	   "meta[name=keywords]"
//	   "meta[property=og:type]"
func GetMetaContent(dom *colly.HTMLElement, metaName string) string {
	meta := dom.ChildAttr(metaName, "content")
	return strings.TrimSpace(meta)
}

func HTML2Text(html string) string {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		crawlers.Sugar.Error(err)
		return ""
	}
	return TrimText(document.Selection)
}
