package main

import "strconv"

// Sum is reponsible to sum the matrix elements and return the result as string
func Sum(records [][]string) (string, error) {

	var sum = 0
	for _, row := range records {
		for i := range row {
			numInt, err := strconv.Atoi(row[i])
			if err != nil {
				return "", err
			}
			sum += numInt
		}
	}
	return strconv.Itoa(sum), nil
}
