package myFind

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	flags "github.com/Arclight-V/Golang/tree/main/module02/ex00/internal/app/flagsSearch"
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
	if fi, _ := d.Info(); fi.Mode()&fs.ModeSymlink == fs.ModeSymlink {
		if p, err := filepath.EvalSymlinks(path); err != nil {
			fmt.Printf("%s -> [broken]\n", path)
		} else {
			fmt.Printf("%s -> %s\n", path, p)
		}
	} else {
		fmt.Println(path)
	}
	return nil
}

func (m *MyFind) findOnlySymLink(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if fi, _ := d.Info(); fi.Mode()&fs.ModeSymlink == fs.ModeSymlink {
		if p, err := filepath.EvalSymlinks(path); err != nil {
			fmt.Printf("%s -> [broken]\n", path)
		} else {
			fmt.Printf("%s -> %s\n", path, p)
		}
	}
	return nil
}

func (m *MyFind) Find() {
	var f func(path string, d fs.DirEntry, err error) error

	switch {
	case m.AllFlagFalse():
		f = m.findAll
	case m.OnlyDir():
		f = m.findOnlyDir
	case m.OnlySymLink():
		f = m.findOnlySymLink
	default:
		f = nil
	}

	err := filepath.WalkDir(m.Folder, f)
	if err != nil {
		log.Println(err)
	}
}
