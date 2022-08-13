package main

import (
	"ex02/internal/app/parser"
	"flag"
	"os"
	"testing"
)

func clearCommandline() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestParseArgNotSecondArgv(t *testing.T) {
	// Arrange
	os.Args = []string{"cmd", "--old", "snapshot1.txt", "--new", ""}
	expectedError := "Error! Empty new path"

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}

	clearCommandline()
}

func TestParseArgNotFirstArgv(t *testing.T) {
	// Arrange
	os.Args = []string{"cmd", "--old", "", "--new", "snapshot2.txt"}
	expectedError := "Error! Empty old path"

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}

	clearCommandline()
}

func TestParseArgNotArgv(t *testing.T) {
	// Arrange
	os.Args = []string{"cmd"}
	expectedError := "Error! Empty old path and new path"

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}

	clearCommandline()
}

func TestParseArgBase(t *testing.T) {
	// Arrange
	expectedFirst := "snapshot1.txt"
	expectedSecond := "snapshot2.txt"
	os.Args = []string{"cmd", "--old", "snapshot1.txt", "--new", "snapshot2.txt"}

	//Actual
	actual, err := parser.ParseArgv()

	// Assert

	if (*actual)[0] != expectedFirst && (*actual)[1] != expectedSecond || err != nil {
		t.Errorf("Incorect result. Expected %s, %s", expectedFirst, expectedSecond)
	}

	clearCommandline()

}
