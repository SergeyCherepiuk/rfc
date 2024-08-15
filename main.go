package main

import (
	"fmt"
	"log"

	"github.com/SergeyCherepiuk/rfc/internal/load"
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

	fmt.Println(string(rfc))
}
