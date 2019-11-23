package main

import "testing"

func TestEcho(t *testing.T) {
	records := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	var expectedMatrixStr = "1,2,3\n4,5,6\n7,8,9\n"
	matrixStr := Echo(records)
	if matrixStr != expectedMatrixStr {
		t.Errorf("Echo was incorrect, got: %s, want: %s", matrixStr, expectedMatrixStr)
	}
}
