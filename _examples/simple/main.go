package main

import (
	"fmt"

	gofeed "github.com/sniperkit/gofeed/pkg"
)

var VERSION string

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println("Title: ", feed.Title)
}
