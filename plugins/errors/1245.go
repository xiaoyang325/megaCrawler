package errors

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1245", "莫斯科国立国际关系学院", "https://mgimo.ru")

	engine.SetStartingURLs([]string{})

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

	extractorConfig.Apply(engine)
}
