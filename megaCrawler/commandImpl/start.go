package commandImpl

func Start(id string) {
	status, err := StartWebsite(id)
	if err != nil {
		println("Service not launched or Invalid :" + err.Error())
		return
	}
	println(status.Message)
}
