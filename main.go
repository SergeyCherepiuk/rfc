package main

import (
	"fmt"
	"log"

	"github.com/SergeyCherepiuk/rfc/internal/load"
	"github.com/SergeyCherepiuk/rfc/internal/utils"
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

	allowList := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	words := utils.SplitWithWhiteList(string(rfc), allowList)

	words = utils.FilterRegularWords(words)

	fmt.Printf("%v\nlen: %d\n", words, len(words))
}
