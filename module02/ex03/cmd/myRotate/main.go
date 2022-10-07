package main

import (
	"fmt"

	parser "github.com/Arclight-V/Golang/tree/52-module02ex03-implement-a-structure-for-transmitting-information/module02/ex03/internal/parser"
)

func main() {
	ds := parser.Parse()
	fmt.Printf(ds)
}
