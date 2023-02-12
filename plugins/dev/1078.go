package dev

import (
	"github.com/PuerkitoBio/goquery"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"net/http"
	"strconv"
)

func init() {
	engine := Crawler.Register("1078", "世界公民参与联盟", "https://www.civicus.org")

	extractorConfig := Extractors.Config{
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
		baseUrl := "https://www.civicus.org/index.php?option=com_minitekwall&task=masonry.getContent&widget_id=11&page="
		for i := 0; true; i++ {
			if Crawler.Test != nil && Crawler.Test.Done {
				return
			}

			pageUrl := baseUrl + strconv.Itoa(i)
			resp, err := http.Get(pageUrl)
			if err != nil {
				continue
			}

			dom, err := goquery.NewDocumentFromReader(resp.Body)
			urls := dom.Find(".mnwall-title > a")
			if len(urls.Nodes) == 0 {
				break
			}
			urls.Each(func(i int, selection *goquery.Selection) {
				pageUrl, ok := selection.Attr("href")
				if !ok {
					return
				}
				engine.Visit(pageUrl, Crawler.News)
			})
		}
	})

	extractorConfig.Apply(engine)
}
