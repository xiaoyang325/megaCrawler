package iiss

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
	"strings"
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
	Html      string `json:"Html"`
	ClassName string `json:"className"`
}

type navComponent struct {
	Current struct {
		Text string `json:"Text"`
		Url  string `json:"Url"`
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
			Url  interface{} `json:"Url"`
			Name interface{} `json:"Name"`
		} `json:"Social"`
		Contact  interface{} `json:"Contact"`
		Detail   string      `json:"Detail"`
		JobTitle interface{} `json:"JobTitle"`
	} `json:"Items"`
}

var reactRegex = regexp.MustCompile("componentRenderQueue.push\\(function\\(\\) \\{ReactDOM.render\\(React.createElement\\(Components.(\\w+), (\\{.+})\\), document.getElementById")

func getReactComponentData(dom *colly.HTMLElement) (component string, data string) {
	match := reactRegex.FindStringSubmatch(dom.Text)
	if len(match) < 3 {
		return
	}
	return match[1], match[2]
}

func init() {
	w := Crawler.Register("iiss", "International Institute for Strategic Studies",
		"https://www.iiss.org/")

	w.SetStartingUrls([]string{
		"https://www.iiss.org/sitemap.xml",
	})

	// 访问文章从 sitemap
	w.OnXML(`//loc`, func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "/blogs/") {
			//w.Visit(element.Text, Crawler.Report)
		} else if strings.Contains(element.Text, "/press/") {
			//w.Visit(element.Text, Crawler.News)
		} else if strings.Contains(element.Text, "/publications/") {
			//w.Visit(element.Text, Crawler.Report)
		} else if strings.Contains(element.Text, "/events/") {
			//w.Visit(element.Text, Crawler.Report)
		} else if strings.Contains(element.Text, "/people/") {
			w.Visit(element.Text, Crawler.Expert)
		}
	})

	w.OnHTML(".container--main script", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if !strings.HasPrefix(element.Text, "componentRenderQueue") {
			return
		}
		component, data := getReactComponentData(element)
		switch ctx.PageType {
		case Crawler.Expert:
			switch component {
			case "Intro":
				var c introComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					Crawler.Sugar.Error(err)
				}
				ctx.Name = c.Title
				ctx.Title = c.Intro
			case "Reading":
				var c readingComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					Crawler.Sugar.Error(err)
				}
				ctx.Description += Crawler.HTML2Text(c.Html) + "\n"
			}
		default:
			switch component {
			case "ArticleNav":
				var c navComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					Crawler.Sugar.Error(err)
				}
				ctx.CategoryText = c.Current.Text
				ctx.PublicationTime = c.Current.Date
			case "Intro":
				var c introComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					Crawler.Sugar.Error(err)
				}
				ctx.Title = Crawler.HTML2Text(c.Title)
				ctx.SubTitle = Crawler.HTML2Text(c.Intro)
			case "Reading":
				var c readingComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					Crawler.Sugar.Error(err)
				}
				ctx.Content += Crawler.HTML2Text(c.Html) + "\n"
			case "AuthorInfo":
				var c authorComponent
				err := json.Unmarshal([]byte(data), &c)
				if err != nil {
					Crawler.Sugar.Error(err)
				}
				for _, item := range c.Items {
					ctx.Authors = append(ctx.Authors, item.Name)
				}
			}
		}
	})
}
