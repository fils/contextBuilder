package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
	"github.com/blevesearch/bleve"
	"github.com/deiu/rdf2go"
	"github.com/kazarena/json-gold/ld"
)

type RWGSchemaorg struct {
	Context      string `json:"@context"`
	Type         string `json:"@type"`
	Name         string `json:"name"`
	ContactPoint struct {
		Type        string `json:"@type"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		URL         string `json:"url"`
		ContactType string `json:"contactType"`
	} `json:"contactPoint"`
	URL    string `json:"url"`
	Funder struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"funder"`
	MemberOf struct {
		Type                string `json:"@type"`
		ProgramName         string `json:"programName"`
		HostingOrganization struct {
			Type string `json:"@type"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"hostingOrganization"`
	} `json:"memberOf"`
	PotentialAction []struct {
		Type   string `json:"@type"`
		Target struct {
			Type        string `json:"@type"`
			URLTemplate string `json:"urlTemplate"`
			Description string `json:"description"`
			HTTPMethod  string `json:"httpMethod"`
		} `json:"target"`
	} `json:"potentialAction"`
}

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(goGetSchemaorg))
	queue := f.Start()

	var domains []string
	// domains = readWhiteList("whitelist_localtest.txt")
	domains = readWhiteList("whitelist.txt")

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

	//  Version that looks for class type
	// doc.Find("script").Each(func(i int, s *goquery.Selection) { // ? script[type="application/ld+json"]
	// 	if s.HasClass("cdfregistry") {
	// 		fmt.Printf("%s\n", s.Text()) //  or send off to a scheme.org parser (JSONLD parser)
	// 	}
	// })

	// Version that just looks for script type application/ld+json
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// s.Has()
		val, _ := s.Attr("type")
		if val == "application/ld+json" {
			// fmt.Printf("%s\n", s.Text()) //  or send off to a scheme.org parser (JSONLD parser)

			// fmt.Printf("%s\n", jsonLDToRDF(s.Text())) //  or send off to a scheme.org parser (JSONLD parser)
			// host := strings.Replace(ctx.Cmd.URL().Host, ".", "", -1)
			// writeFile(fmt.Sprintf("./output/%s.nq", host), jsonLDToRDF(s.Text()))

			bleveIndex(s.Text())
			// convert to RDF (n-triples here and print, ready for loading)
		}
	})
}

// open and save to a bleve index
func bleveIndex(jsonld string) {

	fmt.Printf("Indexed opening\n")

	// Attempt to open an existing index
	var index bleve.Index

	index, berr := bleve.Open("./index/rwg.bleve")
	if berr != nil {
		// open a new index
		mapping := bleve.NewIndexMapping() // do I have to declare this out of the if scope?
		index, berr = bleve.New("./index/rwg.bleve", mapping)
		if berr != nil {
			fmt.Printf("Bleve error making index %v \n", berr)
		}
	}

	var record RWGSchemaorg

	err := json.Unmarshal([]byte(jsonld), &record)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON %v \n", err)
	}

	// index some data
	berr = index.Index(record.URL, record)
	fmt.Printf("Indexed item with URL %s\n", record.URL)
	if berr != nil {
		fmt.Printf("Bleve error indexing %v \n", berr)
	}

	index.Close()
	fmt.Printf("Indexed closed\n\n")

}

// TODO this is a test function for rdf2go
// it does nothing right now
func jsonLDToGraph(jsonld string) {

	baseURI := "https://example.org/foo"

	// Create a new graph
	g := rdf2go.NewGraph(baseURI)
	g.Parse(strings.NewReader(jsonld), "application/ld+json")

	// if err != nil {
	// 	// deal with err
	// }
}

func jsonLDToRDF(jsonld string) string {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	triples, err := proc.ToRDF(myInterface, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err.Error()
	}

	return triples.(string)
}

func writeFile(name string, xmldata string) {
	// Create the output file
	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)

	_, err = w.WriteString(xmldata)
	w.Flush()

	if err != nil {
		log.Fatal(err)
	}
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
