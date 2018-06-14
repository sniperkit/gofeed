package main

import (
	"fmt"

	gofeed "github.com/sniperkit/gofeed/pkg"
	rss "github.com/sniperkit/gofeed/pkg/rss"
)

var VERSION string

var (
	feedData = `<rss version="2.0">
<channel>
<managingEditor>Ender Wiggin</managingEditor>
<itunes:author>Valentine Wiggin</itunes:author>
</channel>
</rss>`
)

type MyCustomTranslator struct {
	defaultTranslator *gofeed.DefaultRSSTranslator
}

func NewMyCustomTranslator() *MyCustomTranslator {
	t := &MyCustomTranslator{}

	// We create a DefaultRSSTranslator internally so we can wrap its Translate
	// call since we only want to modify the precedence for a single field.
	t.defaultTranslator = &gofeed.DefaultRSSTranslator{}
	return t
}

func (ct *MyCustomTranslator) Translate(feed interface{}) (*gofeed.Feed, error) {
	rss, found := feed.(*rss.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *rss.Feed")
	}

	f, err := ct.defaultTranslator.Translate(rss)
	if err != nil {
		return nil, err
	}

	if rss.ITunesExt != nil && rss.ITunesExt.Author != "" {
		f.Author = rss.ITunesExt.Author
	} else {
		f.Author = rss.ManagingEditor
	}
	return f
}

func main() {
	fp := gofeed.NewParser()
	fp.RSSTranslator = NewMyCustomTranslator()
	feed, _ := fp.ParseString(feedData)
	fmt.Println("Author: ", feed.Author) // Valentine Wiggin
}
