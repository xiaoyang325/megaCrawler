package main

import (
	"megaCrawler/megaCrawler"
)
import _ "megaCrawler/plugins/test"
import _ "megaCrawler/plugins/iiss"

func main() {
	megaCrawler.Start()
}
