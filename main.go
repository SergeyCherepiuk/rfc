package main

import (
	"fmt"
	"log"

	"github.com/SergeyCherepiuk/rfc/internal/load"
	"github.com/SergeyCherepiuk/rfc/internal/spellcheck"
	"github.com/SergeyCherepiuk/rfc/internal/transform"
)

func main() {
	loader, err := load.NewRfcLoader()
	if err != nil {
		log.Fatal(err)
	}

	rfc, err := loader.Load(783)
	if err != nil {
		log.Fatal(err)
	}

	words := transform.NewPipeline(rfc).
		AddTransformations(transform.DefaultTransformers...).
		Run()

	checker, err := spellcheck.NewDictionaryChecker()
	if err != nil {
		log.Fatal(err)
	}

	for _, word := range words {
		correct, suggestion, err := checker.Check(word)
		if err != nil {
			log.Fatal()
		}

		if !correct {
			fmt.Printf("%s -> %v\n", word, suggestion)
		}
	}
}
