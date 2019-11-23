package main

import (
	"fmt"
	"strings"
)

func Flatten(records [][]string) string {

	var firstRow = false
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s%s", response, addComma(firstRow), strings.Join(row, ","))
		firstRow = true
	}
	return response
}

func addComma(firstRow bool) string {
	if !firstRow {
		return ""
	} else {
		return ","
	}
}
