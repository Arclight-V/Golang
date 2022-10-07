package main

import (
	"bufio"
	"fmt"
	"os"

	WC "github.com/Arclight-V/Golang/tree/main/module02/ex01/internal/app/myWC"
	parser "github.com/Arclight-V/Golang/tree/main/module02/ex01/internal/app/parser"
)

func scan() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Эхо: %s\n", txt)
	}
}

func main() {
	flagStruct, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	myWC := WC.NewMyWC(flagStruct)
	myWC.Analisys()

}
