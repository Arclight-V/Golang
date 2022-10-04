package tests

import (
	"flag"
	"os"
	"parser"
	"testing"
)

var errorVal = parser.NewErrorsVars()

const (
	appName = "myfind"
	flagF   = "-f"
	flagSL  = "-sl"
	flagD   = "-d"
	flagExt = "-ext"
)

func clearCommandline() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestParseArgvNotArg(t *testing.T) {
	// Arrange
	os.Args = []string{}
	expectedError := errorVal.NothingToLookFor()

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}

}

func TestParseArgvLotsOfFolders(t *testing.T) {
	// Arrange
	os.Args = []string{flagD, "/test", "/test2"}
	expectedError := errorVal.ToManyFolders()
	expectedError1 := errorVal.UnknownFlags()

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || (err.Error() != expectedError && err.Error() != expectedError1) || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}

func TestParseArgvLotsUnknownFlags(t *testing.T) {
	// Arrange
	os.Args = []string{appName, "test", "test2"}
	expectedError := errorVal.ToManyFolders()

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}

	clearCommandline()

}

func TestParseArgvCheckFlagSL(t *testing.T) {
	// Arrange
	os.Args = []string{appName, flagSL}
	expectedError := errorVal.FolderIsNotSpecified()

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}

func TestParseArgvCheckFlagD(t *testing.T) {
	// Arrange
	os.Args = []string{appName, flagD}
	expectedError := errorVal.FolderIsNotSpecified()

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}

func TestParseArgvCheckFlagExt(t *testing.T) {
	// Arrange
	os.Args = []string{appName, flagExt, "test"}
	expectedError := errorVal.OnlyF()

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}

func TestParseArgvCheckFlagF(t *testing.T) {
	// Arrange
	os.Args = []string{appName, flagF}
	expectedError := errorVal.FolderIsNotSpecified()

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}

func TestParseArgvCheckisNoFileAndNoFolder(t *testing.T) {
	// Arrange
	os.Args = []string{appName, "test"}
	_, expectedError := os.Open("test")

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError.Error() || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}

func TestParseArgvCheckisNoFolder(t *testing.T) {
	// Arrange
	os.Args = []string{appName, "test"}
	expectedError := errorVal.IsNotDirectory()
	f, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer os.Remove(os.Args[1])

	//Actual
	actual, err := parser.ParseArgv()

	// Assert
	if err == nil || err.Error() != expectedError || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}

//func TestParseArgNotSecondArgv(t *testing.T) {
//	// Arrange
//	os.Args = []string{"cmd", "--old", "snapshot1.txt", "--new", ""}
//	expectedError := "Error! Empty new path"
//
//	//Actual
//	actual, err := parser.ParseArgv()
//
//	// Assert
//	if err == nil || err.Error() != expectedError || actual != nil {
//		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
//	}
//
//	clearCommandline()
//}
//
//func TestParseArgNotFirstArgv(t *testing.T) {
//	// Arrange
//	os.Args = []string{"cmd", "--old", "", "--new", "snapshot2.txt"}
//	expectedError := "Error! Empty old path"
//
//	//Actual
//	actual, err := parser.ParseArgv()
//
//	// Assert
//	if err == nil || err.Error() != expectedError || actual != nil {
//		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
//	}
//
//	clearCommandline()
//}
//
//func TestParseArgNotArgv(t *testing.T) {
//	// Arrange
//	os.Args = []string{"cmd"}
//	expectedError := "Error! Empty old path and new path"
//
//	//Actual
//	actual, err := parser.ParseArgv()
//
//	// Assert
//	if err == nil || err.Error() != expectedError || actual != nil {
//		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
//	}
//
//	clearCommandline()
//}
//
//func TestParseArgBase(t *testing.T) {
//	// Arrange
//	expectedFirst := "snapshot1.txt"
//	expectedSecond := "snapshot2.txt"
//	os.Args = []string{"cmd", "--old", "snapshot1.txt", "--new", "snapshot2.txt"}
//
//	//Actual
//	actual, err := parser.ParseArgv()
//
//	// Assert
//
//	if (*actual)[0] != expectedFirst && (*actual)[1] != expectedSecond || err != nil {
//		t.Errorf("Incorect result. Expected %s, %s", expectedFirst, expectedSecond)
//	}
//
//	clearCommandline()
//
//}
