import os
import sys

main_template = \
r"""
package dev

import (
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
)

func init() {
	engine := Crawler.Register("%s", "%s", "%s")
	
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
"""

if len(sys.argv) != 4:
    raise RuntimeError("Invalid argument, Usage: python generator.py <website_id> <website_name> <base_url>")

website_id = sys.argv[1]
website_name = sys.argv[2]
base_url = sys.argv[3]

with open("plugins/dev/%s.go" % website_id, "w") as file:
    file.write(main_template % (website_id, website_name, base_url))

print("Generated plugin for %s" % website_name)