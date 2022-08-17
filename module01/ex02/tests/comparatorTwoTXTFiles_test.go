package main

import (
	"comparator"
	"testing"
)

func TestComparatorBase(t *testing.T) {
	// Arrange
	fileName := "snapshot2.txt"
	//Actual										// Assert
	if str, err := comparator.ComparatorTwoTXTFiles(fileName, "../snapshot/snapshot1.txt"); len(str) != 0 || err == nil {
		t.Errorf("Incorect result. Expected error")
	}

	if str, err := comparator.ComparatorTwoTXTFiles("../snapshot/snapshot1.txt", fileName); len(str) != 0 || err == nil {
		t.Errorf("Incorect result. Expected error")
	}

}
