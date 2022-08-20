package main

import (
	"fmt"
	"parser"
)

func main() {

	flagStrut, err := parser.ParseArgv()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*flagStrut)
}
