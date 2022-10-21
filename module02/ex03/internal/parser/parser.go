package parser

import (
	"errors"
	"flag"
	"os"

	dirStruct "github.com/Arclight-V/Golang/tree/main/module02/ex03/internal/dirStruct"
)

func Parse() (*dirStruct.DirStruct, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("Nothing to archive")
	}

	aFlag := flag.String("a", "", "path to archive")
	flag.Parse()
	logs := flag.Args()

	return dirStruct.NewDirStruct(*aFlag, logs), nil
}
