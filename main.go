package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(handler))
	queue := f.Start()
	queue.SendStringGet("http://opencoredata.org", "http://rvdata.us", "http://iedadata.org", "http://bco-dmo.org")
	queue.Close()
}

func handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())

	// doc, err := goquery.NewDocumentFromResponse(res)
	doc, err := goquery.NewDocumentFromResponse(res)

	if err != nil {
		fmt.Printf("[ERR] %s %s - %s\n", ctx.Cmd.Method(), ctx.Cmd.URL(), err)
		return
	}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		fmt.Printf("%s\n", val)
	})

}
