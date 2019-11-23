package main

import "strconv"

func Multiply(records [][]string) (string, error) {

	var product int64 = 1
	for _, row := range records {
		for i := range row {
			numInt, err := strconv.Atoi(row[i])
			if err != nil {
				return "", err
			}
			product *= int64(numInt)
		}
	}
	return strconv.FormatInt(product, 10), nil
}
