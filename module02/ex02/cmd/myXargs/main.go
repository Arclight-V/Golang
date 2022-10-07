package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var str string

	fmt.Scanf("%s\n", &str)

	flag.Parse()
	arg := flag.Args()
	cmdExec := arg[0]
	flags := arg[1:]
	flags = append(flags, str)

	cmd := exec.Command(cmdExec, flags...)
	cmd.Stdin = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
