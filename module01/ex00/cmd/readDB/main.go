package main

import (
	"ex00/internal/app/parser"
	"fmt"
	"handler"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path, err := parser.ParseArgv()
	check(err)

	handler := handler.NewHandler()
	recipes, e := handler.Handle(path)
	check(e)
	fmt.Print(recipes)

}
