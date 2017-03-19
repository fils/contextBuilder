package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(goGetSchemaorg))
	queue := f.Start()

	var domains []string
	domains = readWhiteList("whitelist_localtest.txt")
	// domains = readWhiteList("whitelist.txt")

	queue.SendStringGet(domains...) // note use of variadic parameter on this function
	queue.Close()
}

func readWhiteList(filename string) []string {
	var domains []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return domains
}

func goGetSchemaorg(ctx *fetchbot.Context, res *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())

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

// Walker style handler
func templateHandler(ctx *fetchbot.Context, res *http.Response, err error) {
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
