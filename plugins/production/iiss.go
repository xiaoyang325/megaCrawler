package production

import (
	"encoding/json"
	"regexp"
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/PuerkitoBio/goquery"

	"github.com/gocolly/colly/v2"
)

type introComponent struct {
	Title      string `json:"Title"`
	Stick      bool   `json:"Stick"`
	Intro      string `json:"Intro"`
	SubHeading string `json:"SubHeading"`
	LessText   string `json:"LessText"`
	MoreText   string `json:"MoreText"`
}

type readingComponent struct {
	HTML      string `json:"html"`
	ClassName string `json:"className"`
}

type navComponent struct {
	Current struct {
		Text string `json:"Text"`
		URL  string `json:"URL"`
		Date string `json:"Date"`
	} `json:"Current"`
}

type authorComponent struct {
	HideMobile bool   `json:"HideMobile"`
	Title      string `json:"Title"`
	Heading    string `json:"Heading"`
	Items      []struct {
		Image     string `json:"Image"`
		Name      string `json:"Name"`
		About     string `json:"About"`
		AboutLink string `json:"AboutLink"`
		Social    struct {
			URL  interface{} `json:"URL"`
			Name interface{} `json:"Name"`
		} `json:"Social"`
		Contact  interface{} `json:"Contact"`
		Detail   string      `json:"Detail"`
		JobTitle interface{} `json:"JobTitle"`
	} `json:"Items"`
}

var reactRegex = regexp.MustCompile(`componentRenderQueue.push\(function\(\) \{ReactDOM.render\(React.createElement\(Components.(\w+), (\{.+})\), document.getElementByID`)

func getReactComponentData(dom *colly.HTMLElement) (component string, data string) {
	match := reactRegex.FindStringSubmatch(dom.Text)
	if len(match) < 3 {
		return
	}
	return match[1], match[2]
}

func init() {
	w := crawlers.Register("iiss", "International Institute for Strategic Studies",
		"https://www.iiss.org/")

	w.SetStartingURLs([]string{
		"https://www.iiss.org/sitemap.xml",
	})

	// 访问文章从 sitemap
	w.OnXML(`//loc`, func(element *colly.XMLElement, ctx *crawlers.Context) {
		switch {
		case strings.Contains(element.Text, "/blogs/"):
			w.Visit(element.Text, crawlers.Report)
		case strings.Contains(element.Text, "/press/"):
			w.Visit(element.Text, crawlers.News)
		case strings.Contains(element.Text, "/publications/"):
			w.Visit(element.Text, crawlers.Report)
		case strings.Contains(element.Text, "/events/"):
			w.Visit(element.Text, crawlers.Report)
		case strings.Contains(element.Text, "/people/"):
			w.Visit(element.Text, crawlers.Expert)
		}
	})

	w.OnHTML(".container--main script", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if !strings.HasPrefix(element.Text, "componentRenderQueue") {
			return
		}
		component, data := getReactComponentData(element)
		switch ctx.PageType {
		case crawlers.Expert:
			switch component {
			case "Intro":
				var c introComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					crawlers.Sugar.Error(err)
				}
				ctx.Name = c.Title
				ctx.Title = c.Intro
			case "Reading":
				var c readingComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					crawlers.Sugar.Error(err)
					return
				}
				document, err := goquery.NewDocumentFromReader(strings.NewReader(c.HTML))
				if err != nil {
					crawlers.Sugar.Error(err)
					return
				}
				ctx.Description += extractors.TrimText(document.Selection) + "\n"
			}
		case crawlers.Index, crawlers.News, crawlers.Report:
			switch component {
			case "ArticleNav":
				var c navComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					crawlers.Sugar.Error(err)
				}
				ctx.CategoryText = c.Current.Text
				ctx.PublicationTime = c.Current.Date
			case "Intro":
				var c introComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					crawlers.Sugar.Error(err)
				}
				document, err := goquery.NewDocumentFromReader(strings.NewReader(c.Title))
				if err != nil {
					crawlers.Sugar.Error(err)
					return
				}
				ctx.Title = extractors.TrimText(document.Selection)

				document, err = goquery.NewDocumentFromReader(strings.NewReader(c.Intro))
				if err != nil {
					crawlers.Sugar.Error(err)
					return
				}
				ctx.SubTitle = extractors.TrimText(document.Selection)

			case "Reading":
				var c readingComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					crawlers.Sugar.Error(err)
				}
				document, err := goquery.NewDocumentFromReader(strings.NewReader(c.HTML))
				if err != nil {
					crawlers.Sugar.Error(err)
					return
				}
				ctx.Content += extractors.TrimText(document.Selection) + "\n"
			case "AuthorInfo":
				var c authorComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					crawlers.Sugar.Error(err)
				}
				for _, item := range c.Items {
					ctx.Authors = append(ctx.Authors, item.Name)
				}
			}
		}
	})
}
