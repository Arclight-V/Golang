package main

import (
	handler2 "ex00/internal/app/handler"
	"ex00/internal/app/parser"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path, err := parser.ParseArgv()
	check(err)

	handler := handler2.NewHandler()
	recipes, e := handler.Handle(path)
	check(e)
	fmt.Print(recipes)

}
