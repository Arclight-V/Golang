package main

import (
	"fmt"

	parser "github.com/Arclight-V/Golang/tree/main/module02/ex03/internal/parser"
)

func main() {
	ds, _ := parser.Parse()
	fmt.Println(ds)
}
