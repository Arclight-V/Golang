package main

import (
	"fmt"
	"os"

	WC "github.com/Arclight-V/Golang/tree/main/module02/ex01/internal/app/myWC"
	parser "github.com/Arclight-V/Golang/tree/main/module02/ex01/internal/app/parser"
)

func main() {
	flagStruct, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	myWC := WC.NewMyWC(flagStruct)
	myWC.Analisys()

}
