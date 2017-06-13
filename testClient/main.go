package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

func main() {

	flag.Parse()
	phrase := flag.Arg(0)
	// phrase := "adam sparql"
	fmt.Print(string(callToJSON(phrase)))
}

func callToJSON(phrase string) string {
	indexPath := "/Users/dfils/src/go/src/oceanleadership.org/contextBuilder/index/rwg.bleve"

	index, err := bleve.OpenUsing(indexPath, map[string]interface{}{
		"read_only": true,
	})
	if err != nil {
		log.Printf("error opening index %s: %v", indexPath, err)
	} else {
		log.Printf("registered index: at %s", indexPath)
	}

	// query := bleve.NewMatchQuery(phrase)
	query := bleve.NewQueryStringQuery(phrase)
	search := bleve.NewSearchRequestOptions(query, 10, 0, false) // no explanation
	search.Highlight = bleve.NewHighlight()                      // need Stored and IncludeTermVectors in index
	searchResults, err := index.Search(search)

	hits := searchResults.Hits // array of struct DocumentMatch

	for k, item := range hits {
		fmt.Printf("\n%d: %s, %f, %s, %v\n", k, item.Index, item.Score, item.ID, item.Fragments)
		for key, frag := range item.Fragments {
			fmt.Printf("%s   %s\n", key, frag)
		}
	}

	jsonResults, _ := json.MarshalIndent(hits, " ", " ")

	return string(jsonResults)
}
