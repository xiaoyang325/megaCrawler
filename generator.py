import os
import sys
import re

main_template = \
r"""
package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1001", "欧洲外交关系委员会", "https://ecfr.eu/events/")
	
	engine.SetStartingURLs([]string{"https://ecfr.eu/events/"})
	
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

engine.OnHTML("post-title",func(element *colly.HTMLElement, ctx *crawlers.Context){
    engine.Visit(element.Attr("href"), crawlers.News)
})

engine.OnHTML("page-number page-numbers",func(element *colly.HTMLElement, ctx *crawlers.Context){
    engine.Visit(element.Attr("href"), crawlers.Index)
})

engine.OnHTML("li",func(element *colly.HTMLElement, ctx *crawlers.Context){
    engine.Visit(element.Attr("href"), crawlers.Index)
})

script_input = []
if len(sys.argv) == 4:
    website_id = sys.argv[1]
    website_name = sys.argv[2]
    base_url = sys.argv[3]
    script_input = [(website_id, website_name, base_url)]
else:
    while True:
        i = input()
        if i != "":
            website_id, website_name, base_url = i.split("	")
            script_input.append((website_id, website_name, base_url))
        else:
            break

for website_id, website_name, base_url in script_input:
    with open("plugins/dev/%s.go" % website_id, "w", encoding="utf-8") as file:
        file.write(main_template % (website_id, website_name, base_url))
    print("Generated plugin for ID: %s, Name: %s, Base URL: %s" % (website_id, website_name, base_url))
