package errors

import (
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
)

func init() {
	engine := Crawler.Register("1085", "American-Armenian National Security Institute", "https://aansi.org")

	engine.SetStartingUrls([]string{})

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
