package main

import (
	"ex01/internal/app/parser"
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

	for _, str := range *path {
		fmt.Println(str)
	}

	//handler := handler.NewHandler()
	//recipes, e := handler.Handle(path)
	//check(e)
	//fmt.Print(recipes)

}
