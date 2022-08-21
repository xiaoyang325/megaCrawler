package cd

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"io/ioutil"
	"megaCrawler/megaCrawler"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type resp struct {
	NumOfResults int               `json:"num_of_results"`
	Html         string            `json:"html"`
	Ids          map[string]string `json:"ids"`
	Pagination   struct {
		Page        string `json:"page"`
		Total       string `json:"total"`
		MoreResults bool   `json:"more_results"`
		Message     string `json:"message"`
	} `json:"pagination"`
	Filters []interface{} `json:"filters"`
}

func init() {
	s := megaCrawler.Register("cd", "国际关系研究所", "https://www.clingendael.org/").
		SetTimeout(20 * time.Second)

	processHTML := func(element string) {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(element))
		if err != nil {
			_ = megaCrawler.Logger.Error(err)
			return
		}
		var list []megaCrawler.UrlData

		doc.Find(".title-and-date").Each(func(i int, selection *goquery.Selection) {
			a := selection.Find("a")
			urlString, ok := a.Attr("href")
			if !ok {
				return
			}
			date := selection.Find(".date")
			dateString := date.Text()
			datetime, err := time.Parse("2 January 200604:05", dateString)
			if err != nil {
				_ = megaCrawler.Logger.Error(err)
				return
			}
			list = append(list, megaCrawler.UrlData{Url: urlString, LastMod: datetime})
		})
		for _, data := range list {
			s.AddUrl(data.Url, data.LastMod)
		}
	}

	s.OnLaunch(func() {
		k := func(i int) (r resp) {
			formData := url.Values{
				"page":         {strconv.Itoa(i)},
				"content_type": {"publication"},
				"is_eu_form":   {},
				"search":       {},
			}
			client := &http.Client{}

			//Not working, the post data is not a form
			req, err := http.NewRequest("POST", "https://www.clingendael.org/ajax/archive", strings.NewReader(formData.Encode()))
			if err != nil {
				_ = megaCrawler.Logger.Error(err)
				return
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			resp, err := client.Do(req)
			if err != nil {
				_ = megaCrawler.Logger.Error(err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				_ = megaCrawler.Logger.Error(err)
				return
			}

			err = json.Unmarshal(body, &r)
			if err != nil {
				_ = megaCrawler.Logger.Error(err)
			}
			return
		}

		response := k(1)
		total, err := strconv.Atoi(response.Pagination.Total)
		if err != nil {
			_ = megaCrawler.Logger.Error(err)
			return
		}
		processHTML(response.Html)
		for i := 2; i <= total/response.NumOfResults; i++ {
			i := i
			go func() {
				response := k(i)
				processHTML(response.Html)
			}()
		}
	})

	s.OnHTML(".field--name-body", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})
}
