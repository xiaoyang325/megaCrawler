package extractors

import (
	"megaCrawler/crawlers"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly/v2"
)

func splitTitle(title string, splitter string, hint string) string {
	var largeTextLength int
	var largeTextIndex int
	titlePieces := strings.Split(title, splitter)

	if hint != "" {
		filterRegex := regexp.MustCompile(`[^a-zA-Z\d ]`)
		hint = strings.ToLower(filterRegex.ReplaceAllString(hint, ""))
	}

	for i, piece := range titlePieces {
		filterRegex := regexp.MustCompile(`[^a-zA-Z\d ]`)
		current := strings.TrimSpace(piece)
		if hint != "" && strings.Contains(hint, strings.ToLower(filterRegex.ReplaceAllString(current, ""))) {
			largeTextIndex = i
			break
		}
		if len(current) > largeTextLength {
			largeTextLength = len(current)
			largeTextIndex = i
		}
	}

	title = titlePieces[largeTextIndex]
	return strings.TrimSpace(strings.ReplaceAll(title, "&raquo;", "»"))
}

// getTitle Fetch the article title and analyze it.
//
//	Assumptions:
//	- title tag is the most reliable (inherited from Goose)
//	- h1, if properly detected, is the best (visible to users)
//	- og:title and h1 can help improve the title extraction
func getTitle(dom *colly.HTMLElement) (title string) {
	var filterRegex = regexp.MustCompile("[^\u4e00-\u9fa5a-zA-Z\\d ]")
	title = dom.ChildText("title")
	var useDelimiter bool
	if len(title) == 0 {
		return ""
	}

	titleH1 := ""
	titleH1Slices := dom.ChildTexts("h1")

	if len(titleH1Slices) > 0 {
		sort.Slice(titleH1Slices, func(i, j int) bool {
			return len(titleH1Slices[i]) > len(titleH1Slices[j])
		})
		titleH1 = titleH1Slices[0]
		if len(strings.Split(titleH1, " ")) <= 2 {
			titleH1 = ""
		}
		titleH1 = crawlers.StandardizeSpaces(titleH1)
	}

	titleOG := dom.ChildAttr("meta[property=\"og:title\"]", "content")
	if titleOG != "" {
		titleOG = dom.ChildAttr("meta[name=\"og:title\"]", "content")
	}

	filterTitleText := strings.ToLower(filterRegex.ReplaceAllString(title, ""))
	filterTitleH1 := strings.ToLower(filterRegex.ReplaceAllString(titleH1, ""))
	filterTitleOG := strings.ToLower(filterRegex.ReplaceAllString(titleOG, ""))

	switch {
	case titleH1 == title:
		useDelimiter = true
	case filterTitleH1 != "" && (filterTitleH1 == filterTitleOG) || (strings.Contains(filterTitleText, filterTitleH1) && filterTitleOG != "" && strings.Contains(filterTitleText, filterTitleOG) && len(filterTitleH1) > len(titleOG)):
		title = titleH1
		useDelimiter = true
	case filterTitleOG != "" && strings.HasPrefix(filterTitleText, filterTitleOG) && filterTitleOG != filterTitleText:
		title = titleOG
		useDelimiter = true
	}

	for _, delimiter := range []string{"|", " - ", "_", "/", " » "} {
		if !useDelimiter && strings.Contains(title, delimiter) {
			title = splitTitle(title, delimiter, titleH1)
		}
	}

	title = strings.ReplaceAll(title, "&#65533;", "")

	filterTitle := strings.ToLower(filterRegex.ReplaceAllString(title, ""))
	if filterTitleH1 == filterTitle {
		title = titleH1
	}
	return title
}

func Titles(ctx *crawlers.Context, dom *colly.HTMLElement) {
	title := getTitle(dom)
	if ctx.PageType == crawlers.Expert {
		if ctx.Name != "" {
			return
		}
		ctx.Name = title
	} else {
		if ctx.Title != "" {
			return
		}
		ctx.Title = title
	}
}
