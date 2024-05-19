package main

import (
	"fmt"
	"os"

	"github.com/JosephNaberhaus/agnostic/tool/generator/gen"
	"github.com/JosephNaberhaus/agnostic/tool/generator/reader"
)

func handleErr(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)
}

func main() {
	specs, err := reader.FindAllSpecs()
	if err != nil {
		handleErr(err)
	}

	err = gen.WriteAST(specs)
	if err != nil {
		handleErr(err)
	}
}
