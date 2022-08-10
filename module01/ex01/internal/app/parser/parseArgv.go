package parser

import (
	"errors"
	"flag"
)

func ParseArgv() (*[]string, error) {

	pathOld := flag.String("old", "", "the path to the file to read")
	pathNew := flag.String("new", "", "the path to the file to read")
	flag.Parse()

	if len(*pathOld) == 0 && len(*pathNew) == 0 {
		return nil, errors.New("Error! Empty old path and new path")
	} else if len(*pathOld) == 0 {
		return nil, errors.New("Error! Empty old path")
	} else if len(*pathNew) == 0 {
		return nil, errors.New("Error! Empty new path")
	}
	return &[]string{*pathOld, *pathNew}, nil
}
