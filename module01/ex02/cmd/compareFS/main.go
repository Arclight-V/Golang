package main

import (
	"ex02/internal/app/parser"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {

	path, err := parser.ParseArgv()
	check(err)

	fmt.Printf((*path)[0])
	fmt.Printf((*path)[1])

}
