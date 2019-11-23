package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ValidateMatrix trims the elements and perform two validations
// 1) checks if it is a squared matrix
// 2) checks if the matrix contains only integers
func ValidateMatrix(matrix [][]string) error {

	err1 := validateSquaredMatrix(matrix)
	if err1 != nil {
		return err1
	}
	err2 := validateIntegerMatrix(matrix)
	if err2 != nil {
		return err2
	}
	return nil
}

// check if it is a squared matrix
func validateSquaredMatrix(matrix [][]string) error {

	var rowSize = len(matrix)
	var columnSize = len(matrix[0])
	if rowSize != columnSize {
		return fmt.Errorf("The input matrix is not squared, it is %d/%d", rowSize, columnSize)
	} else {
		return nil
	}
}

// check if it is an integer matrix
func validateIntegerMatrix(matrix [][]string) error {

	for i, row := range matrix {
		for j := range row {
			//trim the element before validation
			matrix[i][j] = strings.TrimSpace(matrix[i][j])

			if _, err := strconv.Atoi(matrix[i][j]); err != nil {
				return fmt.Errorf("The element %s is not an integer", matrix[i][j])
			}
		}
	}
	return nil
}
