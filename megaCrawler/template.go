package megaCrawler

import "github.com/gocolly/colly/v2"

type Template struct {
	htmlHandlers     map[string]func(element *colly.HTMLElement, ctx *Context)
	xmlHandlers      map[string]func(element *colly.XMLElement, ctx *Context)
	responseHandlers []func(response *colly.Response, ctx *Context)
}

func (t *Template) OnHTML(selector string, callback func(element *colly.HTMLElement, ctx *Context)) *Template {
	t.htmlHandlers[selector] = callback
	return t
}

func (t *Template) OnXML(selector string, callback func(element *colly.XMLElement, ctx *Context)) *Template {
	t.xmlHandlers[selector] = callback
	return t
}

func (t *Template) OnResponse(callback func(response *colly.Response, ctx *Context)) *Template {
	t.responseHandlers = append(t.responseHandlers, callback)
	return t
}
