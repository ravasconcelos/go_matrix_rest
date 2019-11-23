package main

import (
	"fmt"
	"strings"
)

// Echo just returns the matrix in string format
func Echo(records [][]string) string {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
