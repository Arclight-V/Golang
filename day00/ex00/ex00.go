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

var arguments = map[string]int{
	"Mean":   Mean,
	"Median": Median,
	"Mod":    Mod,
	"SD":     SD,
}

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

	if buf.Len() != 0 {
		fmt.Print(buf.String())
	}

	if len(m) == 0 {
		return m, errors.New("There are no calculation parameters available")
	}
	return m, nil
}

func calc(metrics []int, input []float64) (out map[string]float64) {

	map_size := len(metrics)
	if map_size == 0 {
		map_size = len(arguments)
	}

	output_metrics := make(map[string]float64, map_size)

	for _, m := range metrics {
		switch m {
		case Mean:
			sum := 0.0
			for _, m := range input {
				sum += m
			}
			output_metrics["Mean"] = sum / float64(len(input))
		case Median:
		case Mod:
		case SD:
		}
	}

	return output_metrics
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
		os.Exit(1)
	}

	fmt.Println("Enter the sequence: \n")
	var input_float []float64
	reader := bufio.NewReader(os.Stdin)
	// TODO:add end-of-file processing and
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

		n, e := strconv.ParseFloat(line, 64)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		if n < -100000 && n > 100000 {
			fmt.Println("Error: Integer numbers between -100000 and 100000.")
			os.Exit(1)
		}

		input_float = append(input_float, n)
	}

	output_metrics := calc(calc_metrics, input_float)

	for key, value := range output_metrics {
		fmt.Printf("%s %c %f\n", key, ':', value)
	}

}
