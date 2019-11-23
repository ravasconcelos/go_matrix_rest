package main

import (
	"fmt"
	"log"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
//      or
//		curl -F 'file=@assets/matrix.csv' "localhost:8080/echo"
// Test with
//		go test
func main() {
	// Change the port if needed
	var HTTP_PORT = "8080"

	//echo: Return the matrix as a string in matrix format
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		HandleRequest(w, r, "echo")
	})

	//invert: Return the matrix as a string in matrix format
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		HandleRequest(w, r, "invert")
	})

	//flatten: Return the matrix as a 1 line string, with values separated by commas
	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		HandleRequest(w, r, "flatten")
	})

	//sum: Return the sum of the integers in the matrix
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		HandleRequest(w, r, "sum")
	})

	//multiply: Return the sum of the integers in the matrix
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		HandleRequest(w, r, "multiply")
	})

	//start the server
	fmt.Printf("Server will be started at localhost: %s\n", HTTP_PORT)
	http.ListenAndServe(":"+HTTP_PORT, nil)
}

// HandleRequest is responsible to logging, validation and calling service method
func HandleRequest(w http.ResponseWriter, r *http.Request, service string) {
	log.Println("Begin handling", service)

	records, err := readCSV(r)
	if err == nil {
		validationErr := ValidateMatrix(records)
		if validationErr == nil {

			switch service {
			case "echo":
				fmt.Fprint(w, Echo(records))
			case "invert":
				fmt.Fprint(w, Invert(records))
			case "flatten":
				fmt.Fprint(w, Flatten(records))
			case "sum":
				sum, err := Sum(records)
				if err != nil {
					w.Write([]byte(fmt.Sprintf("Error in sum: %s", err.Error())))
				} else {
					fmt.Fprint(w, fmt.Sprintf("%s", sum))
				}
			case "multiply":
				product, err := Multiply(records)
				if err != nil {
					w.Write([]byte(fmt.Sprintf("Error in multiplication: %s", err.Error())))
				} else {
					fmt.Fprint(w, fmt.Sprintf("%s", product))
				}
			default:
				w.Write([]byte(fmt.Sprintf("Error: Service %s is not supported", service)))
			}
		} else {
			w.Write([]byte(fmt.Sprintf("Error validating matrix: %s", validationErr.Error())))
		}
	} else {
		w.Write([]byte(fmt.Sprintf("Error reading CSV file: %s", err.Error())))
	}

	log.Println("End handling", service)
}
