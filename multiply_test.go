package main

import "testing"

func TestMultiply3x3(t *testing.T) {
	records := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	var expectedResult = "362880"
	result, err := Multiply(records)
	if err != nil {
		t.Errorf("Error calculating the Multiplication: %s", err.Error())
	} else if result != expectedResult {
		t.Errorf("Multiply was incorrect, got: %s, want: %s", result, expectedResult)
	}
}

func TestMultiply4x4(t *testing.T) {
	records := [][]string{
		{"1", "2", "1", "1"}, // 2
		{"3", "1", "1", "1"}, // 6
		{"1", "1", "4", "1"}, // 24
		{"1", "1", "1", "5"}, // 120
	}
	var expectedResult = "120"
	result, err := Multiply(records)
	if err != nil {
		t.Errorf("Error calculating the Multiplication: %s", err.Error())
	} else if result != expectedResult {
		t.Errorf("Multiply was incorrect, got: %s, want: %s", result, expectedResult)
	}
}

func TestMultiply1x1(t *testing.T) {
	records := [][]string{
		{"9"},
	}
	var expectedResult = "9"
	result, err := Multiply(records)
	if err != nil {
		t.Errorf("Error calculating the Multiplication: %s", err.Error())
	} else if result != expectedResult {
		t.Errorf("Multiply was incorrect, got: %s, want: %s", result, expectedResult)
	}
}

func TestMultiplyInvalidInteger(t *testing.T) {
	records := [][]string{{"1", "X", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	_, err := Multiply(records)
	if err == nil {
		t.Errorf("Multiply error handling was incorrect, got: nil, want: not nil")
	}
}
