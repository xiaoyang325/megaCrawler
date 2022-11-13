package Extractors

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func getMetaImgUrl(dom *colly.HTMLElement) string {
	topMetaImage := GetMetaContent(dom, "meta[property=\"og:image\"]")
	if topMetaImage == "" {
		topMetaImage = dom.ChildAttr("link[rel=\"img_src|image_src\"]", "href")
	}
	if topMetaImage == "" {
		topMetaImage = GetMetaContent(dom, "meta[name=\"og:image\"]")
	}
	if topMetaImage == "" {
		topMetaImage = dom.ChildAttr("link[rel=\"icon\"]", "href")
	}
	if url, err := dom.Request.URL.Parse(topMetaImage); err != nil {
		return url.String()
	} else {
		return topMetaImage
	}
}

func getImgUrls(dom *colly.HTMLElement) (images []string) {
	return dom.ChildAttrs("img", "src")
}

func Image(ctx *Crawler.Context, dom *colly.HTMLElement) {
	ctx.Image = append([]string{getMetaImgUrl(dom)}, getImgUrls(dom)...)
}
