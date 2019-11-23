package main

import (
	"encoding/csv"
	"net/http"
)

func readCSV(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	return records, err
}
