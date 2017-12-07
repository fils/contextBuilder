// +build ignore

package main

import (
	"github.com/kazarena/json-gold/ld"
)

func main() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	triples := `
		<http://example.com/Subj1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://example.com/Type> .
		<http://example.com/Subj1> <http://example.com/prop1> <http://example.com/Obj1> .
		<http://example.com/Subj1> <http://example.com/prop2> "Plain" .
		<http://example.com/Subj1> <http://example.com/prop2> "2012-05-12"^^<http://www.w3.org/2001/XMLSchema#date> .
		<http://example.com/Subj1> <http://example.com/prop2> "English"@en .
	`

	doc, err := proc.FromRDF(triples, options)
	if err != nil {
		panic(err)
	}

	ld.PrintDocument("JSON-LD output", doc)
}
