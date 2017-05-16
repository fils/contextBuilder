package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	ch := make(chan string) // ch := make(chan string)

	sites := []string{"opencoredata.org", "bco-dmo.org", "www.unavco.org", "www.opentopography.org"}
	for _, site := range sites {
		go getJSONLD(site, ch)
	}

	for range sites {
		fmt.Println(<-ch)
	}

}

func getJSONLD(url string, ch chan<- string) {
	resp, err := soup.Get("http://opencoredata.org")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("script", "type", "application/ld+json")

	// need a buffer
	var buffer bytes.Buffer
	for _, link := range links {
		buffer.WriteString(link.Text())
	}
	ch <- buffer.String()
}
