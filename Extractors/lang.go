package Extractors

import (
	"github.com/gocolly/colly/v2"
	"regexp"
)

var ReLang, _ = regexp.Compile("^[A-Za-z]{2}$")

func getMetaLang(dom colly.HTMLElement) string {
	attr := dom.Attr("lang")
	if attr == "" {
		selectors := []string{
			"meta[http-equiv=\"content-language\"]",
			"meta[name=\"lang\"]",
		}
		for _, selector := range selectors {
			meta := dom.ChildAttr(selector, "content")
			if meta != "" {
				attr = meta
				break
			}
		}
	} else {
		value := attr[0:2]
		if ReLang.MatchString(value) {
			return value
		}
	}
	return ""
}
