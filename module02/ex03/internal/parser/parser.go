package parser

import dirStruct "github.com/Arclight-V/Golang/tree/main/module02/ex03/internal/dirStruct"

func Parse() (*dirStruct.DirStruct, error) {
	toArchive := "toArchive"
	toLogs := []string{"logs1", "logs2"}
	return dirStruct.NewDirStruct(toArchive, toLogs), nil
}
