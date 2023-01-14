package theamericanstorypodcast

import (
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("theamericanstorypodcast", "克萊蒙研究所", "https://theamericanstorypodcast.org/")

	w.SetStartingUrls([]string{"https://theamericanstorypodcast.org/", "https://theamericanstorypodcast.org/writings/", "https://theamericanstorypodcast.org/how-to-listen/"})

}
