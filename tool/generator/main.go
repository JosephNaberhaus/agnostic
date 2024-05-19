package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/JosephNaberhaus/agnostic/tool/generator/gen"
	"github.com/JosephNaberhaus/agnostic/tool/generator/model"
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

	slices.SortFunc(specs, func(a, b model.Spec) int {
		return strings.Compare(a.Name, b.Name)
	})

	err = gen.WriteAST(specs)
	if err != nil {
		handleErr(err)
	}

	err = gen.WriteCode(specs)
	if err != nil {
		handleErr(err)
	}
}
