package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	engine := crawlers.Register("1078", "世界公民参与联盟", "https://www.civicus.org")

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

	engine.OnLaunch(func() {
		baseURL := "https://www.civicus.org/index.php?option=com_minitekwall&task=masonry.getContent&widget_id=11&page="
		for i := 0; true; i++ {
			if engine.Test != nil && engine.Test.Done {
				return
			}

			pageURL := baseURL + strconv.Itoa(i)
			resp, err := http.Get(pageURL)
			if err != nil {
				continue
			}

			dom, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				continue
			}
			urls := dom.Find(".mnwall-title > a")
			if len(urls.Nodes) == 0 {
				break
			}
			urls.Each(func(i int, selection *goquery.Selection) {
				pageURL, ok := selection.Attr("href")
				if !ok {
					return
				}
				engine.Visit(pageURL, crawlers.News)
			})

			err = resp.Body.Close()
			if err != nil {
				continue
			}
		}
	})

	extractorConfig.Apply(engine)
}
