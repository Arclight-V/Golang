package parser

import (
	"errors"
	"flag"
	"os"
)

type Flags struct {
	IsSymlinks, IsDirectories, IsFile bool
	FileExtension, Folder             string
}

func ParseArgv() (*Flags, error) {

	if len(os.Args) < 2 {
		return nil, errors.New("nothing to look for")
	}

	isSymlinks := flag.Bool("sl", false, "to print the only symlinks")
	isDirectories := flag.Bool("d", false, "to print the only directories")
	isFile := flag.Bool("f", false, "to print the only files")
	fileExtension := flag.String("ext", "", "which files to search for with the extension")

	flag.Parse()

	if len(*fileExtension) != 0 && !*isFile {
		return nil, errors.New("works ONLY when -f is specified")
	}

	args := flag.Args()
	if len(args) > 1 {
		return nil, errors.New("error! too many arguments")
	}

	allFlag := &Flags{
		IsSymlinks:    *isSymlinks,
		IsDirectories: *isDirectories,
		IsFile:        *isFile,
		FileExtension: *fileExtension,
		Folder:        args[0],
	}

	return allFlag, nil

}
