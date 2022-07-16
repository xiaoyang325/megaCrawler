package commandImpl

import (
	"github.com/schollz/progressbar/v3"
	"time"
)

func Get(id string) {
	website, err := GetWebsite(id)
	if err != nil {
		println("Service not launched or Invalid :" + err.Error())
		return
	}
	println("Website  id:", website.Id)
	println("Website Url:", website.BaseUrl)
	println("Running?   :", website.IsRunning)
	println("Next Run at:", website.NextIter.Format(time.RFC3339))
	if website.IsRunning {
		bar := progressbar.Default(website.TotalUrl)
		bar.Add64(website.DoneUrl)
	}
	println(website.TotalUrl, website.DoneUrl)
	return
}
