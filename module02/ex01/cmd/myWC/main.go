package main

import (
	"fmt"
	"os"

	parser "github.com/Arclight-V/Golang/tree/38-module02ex01-make-base-app-with-pars-flags-in-37-issue/module02/ex01/internal/app/parser"
)

func main() {
	flagStruct, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(flagStruct)

}
