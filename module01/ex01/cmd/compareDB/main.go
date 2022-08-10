package main

import (
	"ex01/internal/app/comparer"
	"ex01/internal/app/parser"
	"fmt"
	"handler"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	path, err := parser.ParseArgv()
	check(err)

	handler := handler.NewHandler()
	recipesOld, e := handler.Handle((*path)[0])
	check(e)

	recipesNew, e := handler.Handle((*path)[1])
	check(e)

	str, ok := comparer.Comparer(recipesOld, recipesNew)
	if !ok {
		fmt.Print(str)
	}
}
