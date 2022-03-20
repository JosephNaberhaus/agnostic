package main

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/implementations"
)

// Generates and runs the tests
func main() {
	for _, language := range implementations.Languages {
		fmt.Printf("Writing tests for %s\n", language)
		err := implementations.WriteAllTests(language)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Running tests for %s\n", language)
		err = implementations.RunAllTests(language)
		if err != nil {
			panic(err)
		}
	}
}
