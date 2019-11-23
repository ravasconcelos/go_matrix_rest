package main

import (
	"testing"
)

func TestSuccessValidateMatrix(t *testing.T) {
	err := ValidateMatrix(createMatrix(10, 10))
	if err != nil {
		t.Errorf("Matrix was not validated properly, got: err, want: nil")
	}
}

func TestValidateSquaredMatrix(t *testing.T) {
	err1 := ValidateMatrix(createMatrix(2, 3))
	if err1 == nil {
		t.Errorf("Matrix was not validated properly, got: nil, want: not nil")
	}

	err2 := ValidateMatrix(createMatrix(3, 2))
	if err2 == nil {
		t.Errorf("Matrix was not validated properly, got: nil, want: not nil")
	}
}

func TestValidateIntegerMatrix(t *testing.T) {
	matrix := createMatrix(3, 3)

	matrix[1][1] = " "
	err1 := ValidateMatrix(matrix)
	if err1 == nil {
		t.Errorf("Matrix was not validated properly, got: nil, want: not nil")
	}

	matrix[1][1] = "1.2"
	err2 := ValidateMatrix(matrix)
	if err2 == nil {
		t.Errorf("Matrix was not validated properly, got: nil, want: not nil")
	}
}

func createMatrix(x int, y int) [][]string {
	matrix := make([][]string, x)
	for i := 0; i < x; i++ {
		matrix[i] = make([]string, y)
		for j := 0; j < y; j++ {
			matrix[i][j] = " 1 "
		}
	}
	return matrix
}
