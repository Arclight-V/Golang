package tests

import (
	"flag"
	parser "github.com/Arclight-V/Golang/tree/main/module02/ex03/internal/parser"
	"os"
	"testing"
)

const (
	appName       = "myRotate"
	flagA         = "-a"
	log1          = "tests/testDate/testLog1.txt"
	log2          = "tests/testDate/testLog2.txt"
	log3          = "tests/testDate/testLog3.txt"
	archiveFolder = "tests/testDate/archiveFolder/"
)

func clearCommandline() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestParseIsFlagIsArchiveIs(t *testing.T) {
	// Arrange
	os.Args = []string{appName, flagA, archiveFolder, log1, log2, log3}

	//Actual
	_, err := parser.Parse()

	// Assert
	if err != nil {
		t.Errorf("\nIncorect result: \n|%s|\nExpected \n|%s|", err, "nil")
	}

	clearCommandline()

}
