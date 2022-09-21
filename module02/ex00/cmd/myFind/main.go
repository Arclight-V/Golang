package main

import (
	"fmt"
	myfind "github.com/Arclight-V/Golang/tree/main/module02/ex00/internal/app/myFind"
	parser "github.com/Arclight-V/Golang/tree/main/module02/ex00/internal/app/parser"
	"os"
)

func main() {

	flagStrut, err := parser.ParseArgv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	mYfind := myfind.NewMyfind(flagStrut)
	mYfind.Find()

}
