package commandImpl

import "time"

func Get(id string) {
	website, err := GetWebsite(id)
	if err != nil {
		println("Service not launched or Invalid :" + err.Error())
		return
	}
	println("Website  id:", website.Id)
	println("Website Url:", website.BaseUrl)
	println("Next Run at:", website.NextIter.Format(time.RFC3339))
	return
}
