package main

import (
	"fmt"
	"os"
	"parser"
)

func main() {

	flagStrut, err := parser.ParseArgv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(*flagStrut)
}
