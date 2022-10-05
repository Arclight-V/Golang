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

func (m *MyFind) printIsDir(path *string, d fs.DirEntry) {
	if d.IsDir() {
		fmt.Println(*path)
	}
}

func (m *MyFind) printIsSymLink(path *string, d fs.DirEntry) bool {
	if fi, _ := d.Info(); fi.Mode()&fs.ModeSymlink == fs.ModeSymlink {
		if p, err := filepath.EvalSymlinks(*path); err != nil {
			fmt.Printf("%s -> [broken]\n", *path)
		} else {
			fmt.Printf("%s -> %s\n", *path, p)
		}
		return true
	}
	return false
}

func (m *MyFind) printFile(path *string) {
	if ext := filepath.Ext(*path); len(ext) != 0 {
		fmt.Println(*path)
	}
}

func (m *MyFind) printFileExt(path *string) {
	if ext := filepath.Ext(*path); ext == m.FileExtension {
		fmt.Println(*path)
	}
}

func (m *MyFind) findOnlyDir(path string, d fs.DirEntry, err error) error {
	if path == m.Folder {
		return nil
	}
	if err != nil {
		return err
	}
	m.printIsDir(&path, d)
	return nil
}

func (m *MyFind) findAll(path string, d fs.DirEntry, err error) error {
	if path == m.Folder {
		return nil
	}
	if err != nil {
		return err
	}
	if m.printIsSymLink(&path, d) {
	} else {
		fmt.Println(path)
	}
	return nil
}

func (m *MyFind) findOnlySymLink(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	m.printIsSymLink(&path, d)
	return nil
}

func (m *MyFind) findOnlyFiles(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	m.printFile(&path)
	return nil
}

func (m *MyFind) findOnlyFilesExt(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	m.printFileExt(&path)
	return nil
}

func (m *MyFind) findMultiple(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if m.IsDirectories {
		m.printIsDir(&path, d)
	}
	if m.IsSymlinks {
		m.printIsSymLink(&path, d)
	}
	if m.IsFile && len(m.FileExtension) == 0 {
		m.printFile(&path)
	}
	if m.IsFile && len(m.FileExtension) > 0 {
		m.printFileExt(&path)
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
	case m.OnlyFiles():
		f = m.findOnlyFiles
	case m.OnlyFilesExt():
		f = m.findOnlyFilesExt
	default:
		f = m.findMultiple
	}

	err := filepath.WalkDir(m.Folder, f)
	if err != nil {
		log.Println(err)
	}
}
