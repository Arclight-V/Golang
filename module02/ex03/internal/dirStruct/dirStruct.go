package dirStruct

import "fmt"

type DirSruct struct {
	pathToArchive string
	pathToLogs    []string
}

func NewDirStruct(toArchive string, toLogs []string) {
	return &DirSruct{pathToArchive: toArchive, pathToLogs: toLogs}
}

func (d *DirSruct) String() {
	fmt.Printf("path to archive: %s\n path to logs: %s", d.pathToArchive, d.pathToLogs)
}
