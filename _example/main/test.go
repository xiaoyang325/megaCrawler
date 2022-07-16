package main

import "megaCrawler"
import _ "megaCrawler/_example/main/plugins/test"
import _ "megaCrawler/_example/main/plugins/test2"

func main() {
	megaCrawler.Start()
}
