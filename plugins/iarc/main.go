package iarc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	s := megaCrawler.Register("iarc", "北极研究所", "https://uaf-iarc.org/").
		SetStartingUrls([]string{"https://uaf-iarc.org/news/page/1/", "https://uaf-iarc.org/news/page/2/"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".fl-post-feed-text", func(element *colly.HTMLElement) {
		t, err := time.Parse("January 2, 2006", element.ChildText(".fl-post-feed-date"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(element.ChildAttr(".fl-post-feed-title > a", "href"), t)
	})

	s.OnResponse(func(response *colly.Response) {
		if strings.Contains(string(response.Body), "Sorry, we couldn't find any posts.") {
			return
		}
		if response.StatusCode == 200 && strings.Contains(response.Request.URL.String(), "page") {
			m1 := regexp.MustCompile(`\d+`)
			num := m1.FindString(response.Request.URL.String())
			k, _ := strconv.Atoi(num)
			s.AddUrl(m1.ReplaceAllString(response.Request.URL.String(), strconv.Itoa(k+1)), time.Now())
		}
	})

	s.OnHTML(".fl-post-title", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})

	s.OnHTML(".fl-post-content", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".fl-post-author", func(element *colly.HTMLElement) {
		megaCrawler.SetAuthor(element, element.ChildText("a"))
	})
}
