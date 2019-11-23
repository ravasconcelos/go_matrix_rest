package main

import "testing"

func TestInvert3x3(t *testing.T) {
	records := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	var expectedMatrixStr = "1,4,7\n2,5,8\n3,6,9\n"
	matrixStr := Invert(records)
	if matrixStr != expectedMatrixStr {
		t.Errorf("Invert was incorrect, got: %s, want: %s", matrixStr, expectedMatrixStr)
	}
}

func TestInvert1x1(t *testing.T) {
	records := [][]string{{"1"}}
	var expectedMatrixStr = "1\n"
	matrixStr := Invert(records)
	if matrixStr != expectedMatrixStr {
		t.Errorf("Invert was incorrect, got: %s, want: %s", matrixStr, expectedMatrixStr)
	}
}

func TestInvert6x6(t *testing.T) {
	records := [][]string{
		{"1", "2", "3", "1", "2", "3"},
		{"4", "5", "6", "4", "5", "6"},
		{"7", "8", "9", "7", "8", "9"},
		{"1", "2", "3", "1", "2", "3"},
		{"4", "5", "6", "4", "5", "6"},
		{"7", "8", "9", "7", "8", "9"},
	}
	var expectedMatrixStr = "1,4,7,1,4,7\n2,5,8,2,5,8\n3,6,9,3,6,9\n1,4,7,1,4,7\n2,5,8,2,5,8\n3,6,9,3,6,9\n"
	matrixStr := Invert(records)
	if matrixStr != expectedMatrixStr {
		t.Errorf("Invert was incorrect, got: %s, want: %s", matrixStr, expectedMatrixStr)
	}
}
