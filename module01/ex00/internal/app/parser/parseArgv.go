package parser

import (
	"errors"
	"flag"
)

func ParseArgv() (p string, err error) {
	path := flag.String("f", "", "the path to the file to read")
	flag.Parse()
	if len(*path) == 0 {
		return "", errors.New("Error! Empty path")
	}
	return *path, nil
}
