package cigi

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strconv"
	"time"
)

type response struct {
	Meta struct {
		TotalCount   int `json:"total_count"`
		Aggregations struct {
			Contentsubtypes   map[string]int `json:"contentsubtypes"`
			TopicsContentpage map[int]int    `json:"topics_contentpage"`
			EventAccess       map[int]int    `json:"event_access"`
			ContentTypes      map[string]int `json:"content_types"`
			TopicsPersonpage  map[int]int    `json:"topics_personpage"`
			Experts           map[int]int    `json:"experts"`
			Years             map[int]int    `json:"years"`
			Contenttypes      map[string]int `json:"contenttypes"`
		} `json:"aggregations"`
	} `json:"meta"`
	Items []struct {
		Elevated       bool      `json:"elevated"`
		Id             int       `json:"id"`
		Title          string    `json:"title"`
		Url            string    `json:"url"`
		Snippet        string    `json:"snippet"`
		PublishingDate time.Time `json:"publishing_date"`
	} `json:"items"`
}

func init() {
	f := false
	s := megaCrawler.Register("cigi", "加拿大国际治理创新中心", "https://www.cigionline.org/").
		SetStartingUrls([]string{"https://www.cigionline.org/api/search/?limit=1&offset=0"})

	s.OnResponse(func(r *colly.Response) {
		if r.Headers.Get("content-type") != "application/json" {
			return
		}
		var re response
		err := json.Unmarshal(r.Body, &re)
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error getting ajax for cigi: %s", err.Error())
			return
		}

		if !f {
			for _, item := range re.Items {
				s.AddUrl(item.Url, item.PublishingDate)
			}
			return
		}
		f = true

		count := re.Meta.TotalCount
		for i := 0; i < count%1000; i++ {
			value, err := s.BaseUrl.Parse("/api/search/")
			if err != nil {
				_ = megaCrawler.Logger.Errorf("Error getting ajax for cigi: %s", err.Error())
				return
			}
			q := value.Query()
			q.Add("limit", "1000")
			q.Add("offset", strconv.Itoa(i*1000))
			q.Add("sort", "date")
			q.Add("field", "publishing_date")
			s.AddUrl(q.Encode(), time.Now())
		}
	})

	s.OnHTML("h1", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})

	s.OnHTML(".col-md-10", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.ChildText("p"))
	})

	s.OnHTML(".block-author", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})
}
