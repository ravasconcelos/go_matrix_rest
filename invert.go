package main

import (
	"fmt"
	"strings"
)

// Invert is reponsible to invert the matrix and return the result as string
func Invert(records [][]string) string {

	// create the new empty matrix
	var matrixSize = len(records[0])
	invertedMatrix := make([][]string, matrixSize)
	for i := 0; i < matrixSize; i++ {
		invertedMatrix[i] = make([]string, matrixSize)
	}
	// invert the matrix
	for i, row := range records {
		for j := range row {
			invertedMatrix[j][i] = records[i][j]
		}
	}

	// save as string
	var response string
	for _, row := range invertedMatrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
