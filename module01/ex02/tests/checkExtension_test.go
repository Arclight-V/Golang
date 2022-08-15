package main

import (
	"checkextension"
	"testing"
)

const (
	TXT = ".txt"
)

func TestCheckExtensionBase(t *testing.T) {
	// Arrange
	fileName := "snapshot2.txt"
	fileName1 := "snapshot2.png"
	//Actual
	ok := checkextension.Checkextension(fileName, TXT)

	// Assert

	if !ok {
		t.Errorf("Incorect result. Expected %s", "true")
	}

	//Actual
	ok1 := checkextension.Checkextension(fileName1, TXT)

	// Assert

	if ok1 {
		t.Errorf("Incorect result. Expected %s", "true")
	}

}
