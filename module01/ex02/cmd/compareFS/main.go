package main

import (
	"checkextension"
	"comparator"
	"fmt"
	"os"
	"parser"
)

const (
	TXT = ".txt"
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

	var ok bool
	var e string
	for _, p := range *path {
		if ok = checkextension.Checkextension(p, TXT); !ok {
			if len(e) == 0 {
				e = fmt.Sprintf("Error! Extension \"%s\", expected: %s", p, TXT)
			} else {
				e += fmt.Sprintf("\nError! Extension \"%s\", expected: %s", p, TXT)
			}
		}
	}
	if !ok {
		check(fmt.Errorf(e))
	}

	str, errcomparator := comparator.ComparatorTwoTXTFiles((*path)[0], (*path)[1])
	check(errcomparator)
	if len(str) != 0 {
		fmt.Print(str)
	}

}
