package tests

import analisysflags "github.com/Arclight-V/Golang/tree/main/module02/ex01/internal/app/analisysFlags"

const (
	appName = "myfind"
	flagI   = "-l"
	flagM  = "-m"
	flagW   = "-w"
	text1 = "/testFiles/text1.txt"
	text2 = "/testFiles/text2.txt"
	text3 = "/testFiles/text3.txt"
	args := []string{flagD, "/testFiles/text1.txt", "/testFiles/text2.txt", "/testFiles/text3.txt"}
	filesToAnalis := []string{"/testFiles/text1.txt", "/testFiles/text2.txt", "/testFiles/text3.txt"}
)

func TestMyWCCountLine(t *testing.T) {
	// Arrange
	os.Args = []string{flagI, text1, text2, text3}
	
	flagStruct := analisysFlags.Flags{flag: flagI, analisisFiles: filesToAnalis}
	myWC := WC.NewMyWC(flagStruct)
	myWC.Analisys()

	var str string
	fmt.Scan(&str)

	// Assert
	if err == nil || (err.Error() != expectedError && err.Error() != expectedError1) || actual != nil {
		t.Errorf("Incorect result |%s|\tExpected |%s|", err.Error(), expectedError)
	}
	clearCommandline()
}