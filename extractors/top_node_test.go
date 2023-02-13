package extractors

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestTopNodeExtractor(t *testing.T) {
	link := "https://www.osce.org/court-of-conciliation-and-arbitration/531383"

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
	rg := regexp.MustCompile(`(\r\n?|\n| ){2,}`)
	k := nodeText(node)
	k = rg.ReplaceAllString(k, "$1")
	t.Log(k)
}
