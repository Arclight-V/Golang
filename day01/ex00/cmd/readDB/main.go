package main

import (
	"ex00/service/dbreader"
	"flag"
)

func main() {

	// TODO: move to parseArgv
	path := flag.String("f", "", "the path to the file to read")
	if len(*path) == 0 {
	}

	jsonreader := dbreader.NewJSONReader()
	xmlreader := dbreader.NewXMLReader()

}
