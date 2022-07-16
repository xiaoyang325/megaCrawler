package main

import (
	"megaCrawler/megaCrawler"
)
import _ "megaCrawler/plugins/test"
import _ "megaCrawler/plugins/test2"

func main() {
	megaCrawler.Start()
}
