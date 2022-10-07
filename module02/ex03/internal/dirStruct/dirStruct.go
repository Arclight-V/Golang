package dirStruct

import "fmt"

type DirStruct struct {
	pathToArchive string
	pathToLogs    []string
}

func NewDirStruct(toArchive string, toLogs []string) *DirStruct {
	return &DirStruct{pathToArchive: toArchive, pathToLogs: toLogs}
}

func (d *DirStruct) String() string {
	return fmt.Sprintf("path to archive: %s\npath to logs: %s", d.pathToArchive, d.pathToLogs)
}
