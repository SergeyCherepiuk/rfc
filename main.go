package main

import (
	"fmt"
	"log"

	"github.com/SergeyCherepiuk/rfc/internal/load"
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

	words := transform.NewTransformer(rfc).
		AddTransformations(transform.DefaultPipeline...).
		Transform()

	fmt.Printf("%v\nlen: %d\n", words, len(words))
}
