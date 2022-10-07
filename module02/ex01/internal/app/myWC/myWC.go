package myWC

import (
	"bufio"
	"fmt"
	"os"

	analisysflags "github.com/Arclight-V/Golang/tree/main/module02/ex01/internal/app/analisysFlags"
)

type MyWC struct {
	*analisysflags.Flags
}

func (wc *MyWC) String() string {
	var f rune

	switch {
	case wc.Mode() == analisysflags.L:
		f = 'l'
	case wc.Mode() == analisysflags.M:
		f = 'm'
	case wc.Mode() == analisysflags.W:
		f = 'w'
	default:
	}

	return fmt.Sprintf("mode: %c,\nfiles: %s", f, wc.Files())
}

func NewMyWC(f *analisysflags.Flags) *MyWC {
	return &MyWC{f}
}

func countLine(fileScanner *bufio.Scanner) int {
	fileScanner.Split(bufio.ScanLines)

	i := 0

	for fileScanner.Scan() {
		fileScanner.Text()
		i++
	}

	return i
}

func countChar(fileScanner *bufio.Scanner) int {
	fileScanner.Split(bufio.ScanLines)

	i := 0

	for fileScanner.Scan() {
		i += len(fileScanner.Text())
	}
	return i
}

func (wc *MyWC) readFile(path string, f func(fileScanner *bufio.Scanner) int) (string, error) {
	readFile, err := os.Open(path)
	defer readFile.Close()

	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(readFile)

	countLine := f(fileScanner)

	return fmt.Sprintf("     %v %s\n", countLine, path), nil
}

func (wc *MyWC) Analisys() {
	var fun func(fileScanner *bufio.Scanner) int

	switch {
	case wc.Mode() == analisysflags.L:
		fun = countLine
	case wc.Mode() == analisysflags.M:
		fun = countChar
	case wc.Mode() == analisysflags.W:
	default:
	}

	analisysFiles := wc.Files()
	type item struct {
		statistic string
		err       error
	}
	ch := make(chan item, wc.LenFiles())
	for _, path := range analisysFiles {
		go func(path string) {
			var it item
			it.statistic, it.err = wc.readFile(path, fun)
			ch <- it
		}(path)
	}

	for i := 0; i < wc.LenFiles(); i++ {
		if it := <-ch; it.err != nil {
			fmt.Printf("error! %s\n", it.err)
		} else {
			fmt.Print(it.statistic)
		}
	}
}
