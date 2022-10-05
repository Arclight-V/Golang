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

func NewMyWC(f *analisysflags.Flags) *MyWC {
	return &MyWC{f}
}

func (wc *MyWC) countLine(path string) (string, error) {
	readFile, err := os.Open(path)
	defer readFile.Close()

	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	i := 0
	for fileScanner.Scan() {
		fileScanner.Text()
		i++
	}

	return fmt.Sprintf("%v   %s\n", i, path), nil
}

func (wc *MyWC) Analisys() {
	var fun func(path string) (string, error)

	switch {
	case wc.Mode() == analisysflags.L:
		fun = wc.countLine
	case wc.Mode() == analisysflags.M:
	case wc.Mode() == analisysflags.W:
	default:
		fmt.Println("test")
	}

	analisysFiles := wc.Files()
	type item struct {
		statistic string
		err       error
	}
	ch := make(chan item, wc.LenFiles())
	for _, f := range analisysFiles {
		go func(f string) {
			var it item
			it.statistic, it.err = fun(f)
			ch <- it
		}(f)
	}

	for i := 0; i < wc.LenFiles(); i++ {
		if it := <-ch; it.err != nil {
			fmt.Printf("error! %s\n", it.err)
		} else {
			fmt.Print(it.statistic)
		}
	}
}
