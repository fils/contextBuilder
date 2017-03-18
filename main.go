package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(goGetSchemaorg))
	queue := f.Start()
	// queue.SendStringGet("http://opencoredata.org", "http://rvdata.us", "http://iedadata.org", "http://bco-dmo.org")
	queue.SendStringGet("http://127.0.0.1:9900/")
	queue.Close()
}

func goGetSchemaorg(ctx *fetchbot.Context, res *http.Response, err error) {
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

	doc.Find("script").Each(func(i int, s *goquery.Selection) { // ? script[type="application/ld+json"]
		if s.HasClass("cdfregistry") {
			fmt.Printf("%s\n", s.Text()) //  or send off to a scheme.org parser (JSONLD parser)
		}
	})

}

// Walker style handerler
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
		fmt.Printf("Found HREF:  %s\n", val)
	})

}

// manifestParser takes the URL of the manifest file and parses it out into the various elements
// If the manigest is JSON-LD read in and convert to RDF and hold as NT to build out a crawl.nt file.
// Look for specific elements in the manifest like: swagger, void, etc and route these to functions for each
// in order to parse and pass them back into the master graph.
func manifestParser() {

}

// Bolt KV function or use RDF and store internal NT file for serilization at end?

// need a function to parse various elements:
// swagger: (might be a lib for that, there are several swagger go packages
// void:  This is RDF..  feed into golang RDF library like I do in other code already
// skosVocs?  (also RDF)
// ontologies?  (RDF, but I think I will pass on this..  too much to try)
