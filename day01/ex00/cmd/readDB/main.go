package main

import (
	"ex00/internal/app/parser"
	"ex00/service/dbreader"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path, err := parser.ParseArgv()
	check(err)

	jsonreader := dbreader.NewJSONReader()
	xmlreader := dbreader.NewXMLReader()

	jsonreader.Read(path)
	xmlreader.Read(path)

}
