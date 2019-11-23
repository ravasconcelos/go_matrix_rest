# Go lang project example for manipulate a Matrix using REST API

In main.go you will find a basic web server written in GoLang. 

The webservice has the ability to perform the following operations:

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo 
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

## Instructions

Prerequisite:
```
    Install go https://golang.org/dl/
```
1. Clone or download this repository 
```
git clone https://github.com/ravasconcelos/go_matrix_rest.git
```
2. Go to the project folder
```
cd go_matrix_rest
```
3. Run the unit tests
```
go test
```
4. Run web server
```
go run . &
```
5. Run the funcional tests
```
cd functional_testing
./send_requests.sh
```
6. Send other requests if you want, just use a valid squared matrix saved in a CSV file
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```
