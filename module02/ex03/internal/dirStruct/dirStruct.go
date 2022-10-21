package dirStruct

import "fmt"

type DirStruct struct {
	pathToArchive string
	pathToLogs    []string
	countLogs     int
}

func NewDirStruct(toArchive string, toLogs []string) *DirStruct {
	return &DirStruct{pathToArchive: toArchive, pathToLogs: toLogs, countLogs: len(toLogs)}
}

func (d *DirStruct) String() string {
	return fmt.Sprintf("path to archive: %s\npath to logs: %s", d.pathToArchive, d.pathToLogs)
}

func (d *DirStruct) PathToLogs() []string {
	return d.pathToLogs
}

func (d *DirStruct) PathToArchive() string {
	return d.pathToArchive
}

func (d *DirStruct) CountLogs() int {
	return d.countLogs
}
