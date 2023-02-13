package commands

import (
	"time"
)

func Get(id string) {
	website, err := GetWebsite(id)
	if err != nil {
		println("Service not launched or Invalid :" + err.Error())
		return
	}
	println("Website Name:", website.Name)
	println("Website  id:", website.Id)
	println("Website Url:", website.BaseUrl)
	println("Running?   :", website.IsRunning)
	println("Next Run at:", website.NextIter.Format(time.RFC3339))
	if website.IsRunning {
		println(website.Bar)
	}
}
