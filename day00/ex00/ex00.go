package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func calc(metrics []int, input []int) {
	for _, m := range metrics {
		switch m {
		case Mean:
		case Median:
		case Mod:
		case SD:
		}
	}
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

	var input_int []int

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		line = strings.TrimSuffix(line, "\n")

		n, e := strconv.Atoi(line)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		if n < -100000 && n > 100000 {
			fmt.Println("Error: Integer numbers between -100000 and 100000.")
			os.Exit(1)
		}

		input_int = append(input_int, n)
	}
	fmt.Print(input_int)
}
