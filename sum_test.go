package main

import "testing"

func TestSum3x3(t *testing.T) {
	records := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	var expectedSum = "45"
	sum, err := Sum(records)
	if err != nil {
		t.Errorf("Error calculating the Sum: %s", err.Error())
	} else if sum != expectedSum {
		t.Errorf("Sum was incorrect, got: %s, want: %s", sum, expectedSum)
	}
}

func TestSum2x2(t *testing.T) {
	records := [][]string{{"1", "2"}, {"4", "5"}}
	var expectedSum = "12"
	sum, err := Sum(records)
	if err != nil {
		t.Errorf("Error calculating the Sum: %s", err.Error())
	} else if sum != expectedSum {
		t.Errorf("Sum was incorrect, got: %s, want: %s", sum, expectedSum)
	}
}

func TestSumInvalidInteger(t *testing.T) {
	records := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "1.2"}}
	_, err := Sum(records)
	if err == nil {
		t.Errorf("Sum error handling was incorrect, got: nil, want: not nil")
	}
}
