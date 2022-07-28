package main

import (
	"fmt"
	"testing"
)

func check(r []int, l []int) bool {
	if len(r) != len(l) {
		fmt.Println("false2")
		return false
	}
	for i, _ := range r {
		if r[i] != l[i] {
			fmt.Printf("r[%v] = %v, l[%v] = %v\n", i, r[i], i, l[i])
			return false
		}
	}
	fmt.Println("true")
	return true

}

func TestParseArg(t *testing.T) {

	// Arrange
	arg0 := []string{"test", "test", "test", "test"}
	arg1 := []string{}

	arg2 := []string{"test", "Mean", "SD", "test"}
	// Act

	metrics0, result0 := parseArg(arg0)
	metrics1, result1 := parseArg(arg1)
	metrics2, result2 := parseArg(arg2)

	// Assert

	if result0 == nil && len(metrics0) != 0 {
		t.Errorf("Incorect result. Expected %s, %s", arg0, result0)
	}

	if result1 != nil && len(metrics1) == 0 {
		t.Errorf("Incorect result. Expected %v, %v", []int{Mean, Median, Mod, SD}, nil)
	}

	if result2 != nil || check(metrics2, []int{Median, SD}) != true {
		t.Errorf("Incorect result. Expected %v, %s", []int{Median, SD}, []string{"test", "test"})
	}

}

func TestCheckCountArg(t *testing.T) {
	// Arrange
	arg0 := []string{}

	arg1 := []string{"test", "test", "test", "test", "test"}
	expected1 := "Exceeded the number of arguments. Maximum 4"

	arg2 := []string{"test", "test", "test", "test"}

	// act
	result0 := checkCountArg(arg0)
	result1 := checkCountArg(arg1)
	result2 := checkCountArg(arg2)

	// Assert

	if result0 != nil {
		t.Errorf("Incorect result. Expected %#v, %s", nil, result0)
	}
	if result1 == nil {
		t.Errorf("Incorect result. Expected %#v, %s", expected1, result1)
	}
	if result2 != nil {
		t.Errorf("Incorect result. Expected %#v, %s", nil, result2)
	}
}
