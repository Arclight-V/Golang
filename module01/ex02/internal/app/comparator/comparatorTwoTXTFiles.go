package comparator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ComparatorTwoTXTFiles(path1 string, path2 string) (string, error) {

	file1, err1 := os.Open(path1)
	if err1 != nil {
		return "", err1
	}
	defer file1.Close()

	scanner := bufio.NewScanner(file1)
	// scanner.Split(bufio.ScanLines)
	linesInFile1 := make(map[string]bool)

	for scanner.Scan() {
		linesInFile1[scanner.Text()] = true
	}

	file2, err2 := os.Open(path2)
	if err2 != nil {
		return "", err2
	}
	defer file2.Close()

	scanner = bufio.NewScanner(file2)

	var sb strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		ok := linesInFile1[line]
		if !ok {
			sb.WriteString(
				fmt.Sprintf("ADDED %s\n", line),
			)
		} else {
			linesInFile1[line] = false
		}
	}

	for k, v := range linesInFile1 {
		if v == true {
			sb.WriteString(
				fmt.Sprintf("REMOVED %s\n", k),
			)
		}
	}

	return sb.String(), nil

}
