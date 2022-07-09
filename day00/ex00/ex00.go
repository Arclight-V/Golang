package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	Mean = iota
	Median
	Mod
	SD
)

func checkCountArg(arg []string) (err error) {
	if len(arg) > 4 {
		return errors.New("Exceeded the number of arguments. Maximum 4")
	}
	return nil
}

func parseArg(arg []string) (metrics []int, err error) {
	if len(arg) == 0 {
		return []int{Mean, Median, Mod, SD}, nil
	}

	arguments := map[string]int{
		"Mean":   Mean,
		"Median": Median,
		"Mod":    Mod,
		"SD":     SD}
	var buf bytes.Buffer
	var m []int

	for c, a := range arg {
		found := false
		for Key, Value := range arguments {
			if Key == a {
				found = true
				m = append(m, Value)
			}
		}
		if found == false {
			if buf.Len() == 0 {
				buf.WriteString("Error: Invalid arguments:\n")
			}
			buf.WriteString(strconv.Itoa(c+1) + " - " + a + "\n")
		}
	}

	return metrics, errors.New(buf.String())
}

func main() {

	arg := os.Args[1:]

	err := checkCountArg(arg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	calc_metrics, err := parseArg(arg)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(calc_metrics)

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(text)
	}
}
