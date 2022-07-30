package parser

import (
	"errors"
	"flag"
)

func ParseArgv() (*[]string, error) {
	fs := flag.NewFlagSet("", flag.ExitOnError)

	pathOld := fs.String("--old", "", "the path to the file to read")
	pathNew := fs.String("--new", "", "the path to the file to read")
	fs.Parse([]string{"--old", "--new"})
	if len(*pathOld) == 0 {
		return nil, errors.New("Error! Empty old path")
	}
	if len(*pathNew) == 0 {
		return nil, errors.New("Error! Empty new path")
	}
	return &[]string{*pathOld, *pathNew}, nil
}
