package production

import (
	"megaCrawler/crawlers"
)

func init() {
	w := crawlers.Register("theamericanstorypodcast", "克萊蒙研究所", "https://theamericanstorypodcast.org/")

	w.SetStartingUrls([]string{"https://theamericanstorypodcast.org/", "https://theamericanstorypodcast.org/writings/", "https://theamericanstorypodcast.org/how-to-listen/"})
}
