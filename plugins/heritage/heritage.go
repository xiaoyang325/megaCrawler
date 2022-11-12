package heritage

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"regexp"
	"strings"
	"time"
)

func init() {
	w := megaCrawler.Register("heritage", "美国传统基金会", "https://www.heritage.org/")
	w.SetStartingUrls([]string{"https://www.heritage.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Text, "?page=") {
			w.Visit(element.Text, megaCrawler.Index)
		} else {
			if strings.Contains(element.Text, "staff") {
				w.Visit(element.Text, megaCrawler.Expert)
			}
			w.Visit(element.Text, megaCrawler.News)
		}
	})

	w.OnHTML(".headline", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.PageType == megaCrawler.Expert {
			ctx.Name = element.Text
		} else {
			ctx.Title = element.Text
		}
	})

	w.OnHTML(".expert-bio-card__expert-title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".expert-card__expert-name", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".author-card__name", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".author-card__multi-name", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".article__body-copy", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.PageType == megaCrawler.Expert {
			ctx.Description = strings.TrimSpace(element.Text)
		} else {
			ctx.Content = strings.TrimSpace(element.Text)
		}
	})

	w.OnHTML("meta[property=\"og:description\"]", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.PageType != megaCrawler.Expert {
			ctx.Description = element.Attr("content")
		}
	})

	w.OnHTML(".expert-bio-card__photo", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		style := element.Attr("style")
		reg, _ := regexp.Compile("https?://(www\\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\\.[a-zA-Z\\d()]{1,6}\\b([-a-zA-Z0-9@:%_+.~#?&/=]*)")
		ctx.Image = []string{reg.FindString(style)}
	})

	w.OnHTML(".article__eyebrow > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})

	w.OnHTML(".article-general-info", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		reg, _ := regexp.Compile("(\\w+ \\d+, \\d+)")
		match := reg.FindString(element.Text)
		times, err := time.Parse("Jan 2, 2006", match)
		if err != nil {
			megaCrawler.Sugar.Errorw(err.Error(), "Original", element.Text, "Regex", match)
			return
		}
		ctx.PublicationTime = times.Format(time.RFC3339)
	})
}
