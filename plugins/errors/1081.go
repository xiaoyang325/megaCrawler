package errors

import (
	"github.com/temoto/robotstxt"
	"io"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"net/http"
)

func init() {
	engine := Crawler.Register("1081", "American Bar Association Center for Global Programs", "https://www.americanbar.org")

	engine.OnLaunch(func() {
		resp, err := http.Get("https://www.americanbar.org/robots.txt")
		if err != nil {
			Crawler.Sugar.Error(err)
			return
		}
		robots, err := robotstxt.FromResponse(resp)
		if err != nil {
			Crawler.Sugar.Error(err)
			return
		}
		err = resp.Body.Close()
		if err != nil {
			Crawler.Sugar.Error(err)
			return
		}
		for _, sitemap := range robots.Sitemaps {
			resp, err := http.Get(sitemap)
			if err != nil {
				Crawler.Sugar.Error(err)
				continue
			}
			read, err := io.ReadAll(resp.Body)
			if err != nil {
				Crawler.Sugar.Error(err)
				continue
			}
			Crawler.Sugar.Infof("%s", string(read))

			//fz, err := gzip.NewReader(resp.Body)
			//if err != nil {
			//	Crawler.Sugar.Error(err)
			//	continue
			//}
			//reader, err := goquery.NewDocumentFromReader(fz)
			//if err != nil {
			//	Crawler.Sugar.Error(err)
			//	return
			//}
			//reader.Find("//loc").Each(func(i int, selection *goquery.Selection) {
			//	url := selection.Text()
			//	if strings.Contains(url, "/news/") || strings.Contains(url, "/groups/") {
			//		engine.Visit(selection.Text(), Crawler.News)
			//	}
			//})
		}
	})

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

	extractorConfig.Apply(engine)
}
