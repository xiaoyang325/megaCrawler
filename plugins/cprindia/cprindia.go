package cprindia

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"fmt"
)

func init() {
	w := megaCrawler.Register("Centre for Policy Research", "政策研究中心", "https://cprindia.org/")

	w.SetStartingUrls([]string{
		"https://cprindia.org/researcharea/agriculture/",
		"https://cprindia.org/researcharea/air-pollution/",
		"https://cprindia.org/researcharea/economy/",
		"https://cprindia.org/researcharea/climate-change/",
		"https://cprindia.org/researcharea/education/",
		"https://cprindia.org/researcharea/environmental-law-justice/",
		"https://cprindia.org/researcharea/energy-electricity/",
		"https://cprindia.org/researcharea/federalism/",
		"https://cprindia.org/researcharea/governance-accountability-public-finance/",
		"https://cprindia.org/researcharea/health-nutrition/",
		"https://cprindia.org/researcharea/indian-politics/",
		"https://cprindia.org/researcharea/international-relations-security/",
		"https://cprindia.org/researcharea/jobs/",
		"https://cprindia.org/researcharea/land-rights/",
		"https://cprindia.org/researcharea/sanitation/",
		"https://cprindia.org/researcharea/social-justice/",
		"https://cprindia.org/researcharea/state-capacity/",
		"https://cprindia.org/researcharea/technology/",
		"https://cprindia.org/researcharea/urbanisation/",
		"https://cprindia.org/researcharea/water/",
		"https://cprindia.org/researcharea/miscellaneous/",
		"https://cprindia.org/research/accountability-initiative/",
		"https://cprindia.org/research/india-infrastructures-and-ecologies-program/",
		"https://cprindia.org/research/initiative-on-cities-economy-and-society/",
		"https://cprindia.org/research/initiative-on-climate-energy-environment/",
		"https://cprindia.org/research/land-rights-initiative/",
		"https://cprindia.org/research/scaling-cities-institutions-for-india/",
		"https://cprindia.org/research/state-capacity-initiative/",
		"https://cprindia.org/research/the-jobs-initiative/",
		"https://cprindia.org/research/the-politics-initiative/",
		"https://cprindia.org/research/the-technology-society-initiative/",
		"https://cprindia.org/research/treads-transboundary-rivers-ecologies-development-studies/",
		"https://cprindia.org/events/",
	})

	w.OnHTML("#tab2>div.tab-news-sec>div.tab-news-text>a",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			fmt.Println(element.Attr("href"))
		})
}