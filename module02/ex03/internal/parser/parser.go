package parser

import dirStruct "github.com/Arclight-V/Golang/tree/52-module02ex03-implement-a-structure-for-transmitting-information/module02/ex03/internal/dirStruct"

func Parse() (*dirStruct.dirStruct, error) {
	toArchive := "toArchive"
	toLogs := []string{"logs1", "logs2"}
	return dirStruct.NewDirStruct{toArchive, toLogs}
}
