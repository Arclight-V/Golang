package main

import (
	"fmt"
	"os"

	parser "github.com/Arclight-V/Golang/tree/main/module02/ex03/internal/parser"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	ds, err := parser.Parse()
	check(err)
	fmt.Println(ds)
}
