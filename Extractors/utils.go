package Extractors

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

type selectorContentPair struct {
	selector string
	content  string
}

// GetMetaContent Extract a given meta content form document.
//	Example metaNames:
//	   "meta[name=description]"
//	   "meta[name=keywords]"
//	   "meta[property=og:type]"
func GetMetaContent(dom *colly.HTMLElement, metaName string) string {
	meta := dom.ChildAttr(metaName, "content")
	return strings.TrimSpace(meta)
}

func ignoreError[T any](s T, err error) T {
	return s
}

func unwrapError[T any](s T, err error) T {
	if err != nil {
		panic(err)
	}
	return s
}
