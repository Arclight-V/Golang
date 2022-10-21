package myRotate

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	dirStruct "github.com/Arclight-V/Golang/tree/main/module02/ex03/internal/dirStruct"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type MyRotate struct {
	*dirStruct.DirStruct
}

type item struct {
	statistic string
	err       error
}

var ch chan item

func NewMyRotate(ds *dirStruct.DirStruct) *MyRotate {
	return &MyRotate{DirStruct: ds}
}

func (mr *MyRotate) archiveLog(pathToArchive, pathToFile string) (string, error) {

	var checkExists = func(toCheck string) bool {
		if _, err := os.Stat(toCheck); os.IsNotExist(err) {
			return false
		}
		return true
	}

	checkExists(pathToArchive)
	if isOk := checkExists(pathToArchive); !isOk {
		return "", errors.New(fmt.Sprintf("<%s> not exists!!!", pathToArchive))
	}
	if isOk := checkExists(pathToFile); !isOk {
		return "", errors.New(fmt.Sprintf("<%s> not exists!!!", pathToArchive))
	}

	f, err := os.Open(pathToFile)
	if err != nil {
		return "", err
	}

	read := bufio.NewReader(f)

	data, errRead := io.ReadAll(read)
	if errRead != nil {
		return "", errRead
	}
	if pathToArchive[len(pathToArchive)-1:] != "/" {
		pathToArchive = pathToArchive + "/"
	}
	nameOfFile := pathToArchive + filepath.Base(f.Name()) + "_" + strconv.Itoa(int(time.Now().Unix())) + ".gz"
	f, errCreat := os.Create(nameOfFile)
	if errCreat != nil {
		return "", errCreat
	}
	w := gzip.NewWriter(f)
	w.Write(data)
	w.Close()

	return nameOfFile, nil
}

func (mr *MyRotate) Rotate() {

	pathToArchive := mr.PathToArchive()
	ch := make(chan item, mr.CountLogs())

	if len(pathToArchive) == 0 {
		for _, log := range mr.PathToLogs() {
			go func(log string) {
				var it item
				it.statistic, it.err = mr.archiveLog(filepath.Dir(log), log)
				ch <- it
			}(log)
		}
	} else {
		for _, log := range mr.PathToLogs() {
			go func(log string) {
				var it item
				it.statistic, it.err = mr.archiveLog(pathToArchive, log)
				ch <- it
			}(log)
		}
	}
	for i := 0; i < mr.CountLogs(); i++ {
		if it := <-ch; it.err != nil {
			fmt.Printf("error! %s\n", it.err)
		} else {
			fmt.Printf("Archive create: %s\n", it.statistic)
		}
	}
}
