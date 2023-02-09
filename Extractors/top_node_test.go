package Extractors

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"testing"
)

func TestTopNodeExtractor(t *testing.T) {
	link := "https://www.valuewalk.com/qualivian-investment-partners-3q22-commentary-floor-decor/"

	resp, err := http.Get(link)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	node := CalculateBestNode(dom.Selection, "en")
	node = postCleanup(node, "en")
	if node == nil {
		t.Error("node is nil")
		return
	}
	rg := regexp.MustCompile(`(\r\n?|\n){2,}`)
	k := nodeText(node)
	println(rg.ReplaceAllString(k, "$1"))
}
