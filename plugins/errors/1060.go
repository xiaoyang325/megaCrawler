package errors

import (
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
)

func init() {
	engine := Crawler.Register("1060", "Kidwai纪念肿瘤研究所", "https://kmio.karnataka.gov.in")

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
