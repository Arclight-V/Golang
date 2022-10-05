package parser

import (
	"errors"
	"flag"
	"os"

	analisysflags "github.com/Arclight-V/Golang/tree/38-module02ex01-make-base-app-with-pars-flags-in-37-issue/module02/ex01/internal/app/analisysFlags"
)

func Parse() (*analisysflags.Flags, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("Nothing analisys")
	}

	boolArr := make([]*bool, 3)

	boolArr[analisysflags.L] = flag.Bool("l", false, "for counting lines")
	boolArr[analisysflags.M] = flag.Bool("m", false, "for counting characters")
	boolArr[analisysflags.W] = flag.Bool("w", false, "for counting words")

	flag.Parse()

	count := 0
	var f analisysflags.FlagMode
	for i, b := range boolArr {
		if *b == true {
			count++
			f = analisysflags.FlagMode(i)
		}
		if count > 1 {
			return nil, errors.New("error! only one flag is possible")
		}
	}

	return &analisysflags.Flags{Flag: f, AnalisisFiles: flag.Args()}, nil
}
