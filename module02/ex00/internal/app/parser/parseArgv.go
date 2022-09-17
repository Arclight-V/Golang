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

type ErrorsVars struct {
	folderIsNotSpecified string
	nothingToLookFor     string
	printSmlinks         string
	printDirectories     string
	printOnlyFiles       string
	searchWithExtension  string
	unknownFlags         string
	onlyF                string
	toManyFolders        string
}

func NewErrorsVars() *ErrorsVars {
	nev := &ErrorsVars{folderIsNotSpecified: "error! the search folder is not specified",
		nothingToLookFor:    "nothing to look for",
		printSmlinks:        "to print the only symlinks",
		printDirectories:    "to print the only directories",
		printOnlyFiles:      "to print the only files",
		searchWithExtension: "which files to search for with the extension",
		unknownFlags:        "error! unknown flags",
		onlyF:               "works ONLY when -f is specified",
		toManyFolders:       "error! too many folders for search",
	}

	return nev
}

// "error! the search folder is not specified"
func (e *ErrorsVars) FolderIsNotSpecified() string {
	return e.folderIsNotSpecified
}

// "nothing to look for"
func (e *ErrorsVars) NothingToLookFor() string {
	return e.nothingToLookFor
}

// "to print the only symlinks"
func (e *ErrorsVars) PrintSymLinks() string {
	return e.printSmlinks
}

// "to print the only directories"
func (e *ErrorsVars) PrintDirectories() string {
	return e.printDirectories
}

// "to print the only files"
func (e *ErrorsVars) PrintOnlyFiles() string {
	return e.printOnlyFiles
}

// "which files to search for with the extension"
func (e *ErrorsVars) SearchWithExtension() string {
	return e.searchWithExtension
}

// "error! unknown flags"
func (e *ErrorsVars) UnknownFlags() string {
	return e.unknownFlags
}

// "works ONLY when -f is specified"
func (e *ErrorsVars) OnlyF() string {
	return e.onlyF
}

// "error! too many folders for search"
func (e *ErrorsVars) ToManyFolders() string {
	return e.toManyFolders
}

func ParseArgv() (*Flags, error) {

	errStruct := NewErrorsVars()

	if len(os.Args) < 2 {
		return nil, errors.New(errStruct.NothingToLookFor())
	}

	isSymlinks := flag.Bool("sl", false, errStruct.PrintSymLinks())
	isDirectories := flag.Bool("d", false, errStruct.PrintDirectories())
	isFile := flag.Bool("f", false, errStruct.PrintOnlyFiles())
	fileExtension := flag.String("ext", "", errStruct.SearchWithExtension())

	flag.Parse()

	if !*isSymlinks && !*isDirectories && !*isFile && len(*fileExtension) == 0 {
		return nil, errors.New(errStruct.UnknownFlags())
	}

	if len(*fileExtension) != 0 && !*isFile {
		return nil, errors.New(errStruct.OnlyF())
	}

	args := flag.Args()
	if len(args) > 1 {
		return nil, errors.New(errStruct.ToManyFolders())
	} else if len(args) == 0 {
		return nil, errors.New(errStruct.FolderIsNotSpecified())
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
