package myFind

import (
	"fmt"
	flags "github.com/Arclight-V/Golang/tree/main/module02/ex00/internal/app/flagsSearch"
	"io/fs"
	"log"
	"path/filepath"
)

type MyFind struct {
	*flags.Flags
}

func NewMyfind(fl *flags.Flags) *MyFind {
	return &MyFind{fl}
}

func (m *MyFind) findOnlyDir(path string, d fs.DirEntry, err error) error {
	if path == m.Folder {
		return nil
	}
	if err != nil {
		return err
	}
	if d.IsDir() {
		fmt.Println(path)
	}
	return nil
}

func (m *MyFind) findAll(path string, d fs.DirEntry, err error) error {
	if path == m.Folder {
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func (m *MyFind) Find() {
	var f func(path string, d fs.DirEntry, err error) error

	if m.AllFlagFalse() {
		f = m.findAll
	} else if m.OnlyDir() {
		f = m.findOnlyDir
	}

	err := filepath.WalkDir(m.Folder, f)
	if err != nil {
		log.Println(err)
	}
}
