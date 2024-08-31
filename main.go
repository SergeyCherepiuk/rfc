package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/SergeyCherepiuk/rfc/internal/load"
	"github.com/SergeyCherepiuk/rfc/internal/pool"
	"github.com/SergeyCherepiuk/rfc/internal/spellcheck"
	"github.com/SergeyCherepiuk/rfc/internal/spellcheck/dictionary"
	"github.com/SergeyCherepiuk/rfc/internal/transform"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	loader, err := load.NewRfcLoader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	checker, err := dictionary.NewDictionaryChecker()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	pool := pool.NewPool(ctx, 5, func(number int) ([]spellcheck.CheckResult, error) {
		return loadAndProcessRfc(loader, checker, number)
	})

	go func() {
		for i := 1; i <= 20; i++ {
			pool.In() <- i
		}
		pool.Close()
	}()

	for {
		select {
		case result, ok := <-pool.Out():
			if !ok {
				return
			}

			for _, r := range result {
				fmt.Printf("%s -> %v\n", r.Word, r.Suggestions)
			}
		case err, ok := <-pool.Err():
			if !ok {
				return
			}

			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
}

func loadAndProcessRfc(
	loader load.RfcLoader,
	checker spellcheck.Checker,
	number int,
) ([]spellcheck.CheckResult, error) {
	log.Printf("Loading and processing RFC %d\n", number)

	rfc, err := loader.Load(number)
	if err != nil {
		return nil, err
	}

	words := transform.NewPipeline(rfc).
		AddTransformations(transform.DefaultTransformers...).
		Run()

	results := make([]spellcheck.CheckResult, 0)

	for result, err := range checker.IncorrectWords(words) {
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
