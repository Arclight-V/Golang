package parser

import (
	"errors"
	"flag"
)

func ParseArgv() (pOld, pNew string, err error) {
	pathOld := flag.String(" --old", "", "the path to the file to read for comparison")
	pathNew := flag.String(" --new", "", "the path to the file to read for comparison")
	flag.Parse()
	if len(*pathOld) == 0 {
		return "", "", errors.New("Error! Empty path")
	}
	if len(*pathNew) == 0 {
		return "", "", errors.New("Error! Empty path")
	}
	return *pathOld, *pathNew, nil
}
