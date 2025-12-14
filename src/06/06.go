package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	// testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	total := 0

	rows := len(data) - 1 // remove one since the last one contains the operations
	columns := len(strings.Fields(data[0]))

	numbers := make([][]int, rows)
	operations := make([]string, columns)

	for i, row := range data {

		// Remove the whitespace and split
		row := strings.Fields(row)

		if i == rows {
			for j, v := range row {
				operations[j] = v
			}
			break
		}

		// Create local row numbers
		rowNumbers := make([]int, columns)
		for j, v := range row {

			intNumber, _ := strconv.Atoi(v)

			rowNumbers[j] = intNumber
		}

		numbers[i] = rowNumbers

	}

	for i := 0; i < columns; i++ {

		localSum := 0
		if operations[i] == "*" {
			localSum = 1
		}

		for j := 0; j < rows; j++ {

			if operations[i] == "*" {
				localSum *= numbers[j][i]
			} else {
				localSum += numbers[j][i]
			}
		}
		total += localSum
	}

	return total
}

func part222(data []string) int {

	numberOfRows := len(data) - 1 // 0 index, last row is operations
	numberOfCols := len(strings.Fields(data[0]))

	operations := make([]string, numberOfCols)

	// First convert each row to a column
	numbers := make([][]string, numberOfCols)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = []string{}
	}

	for i := 0; i <= numberOfRows; i++ {

		// Split multiple whitespaces
		rowArray := strings.Fields(data[i])

		// Last row operations
		if i == numberOfRows {
			copy(operations, rowArray)
			break
		}

		// Rest are the numbers
		// Keep it as string to be able to split later into digits
		for j := 0; j < numberOfCols; j++ {
			numbers[j] = append(numbers[j], rowArray[j])
		}
	}

	total := 0
	for c := 0; c < len(numbers); c++ {
		total += rotateArrayAndCalculate(numbers[c], operations[c])
	}

	return total
}

func rotateArrayAndCalculate(column []string, operation string) int {

	// fmt.Println(column, operation) // DEBUG

	// Step 1 : Create the new 2D array. Max size of number is 4 digits
	numbers := make([][]int, 4)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = make([]int, 4)
	}

	// Step 2 : Split each digit of the integers and put into 2D array :
	for i := 0; i < len(column); i++ {

		digits := []rune(column[i])

		counter := 3
		// Parse the digits from the end to the start
		for j := len(digits) - 1; j >= 0; j-- {
			intDigit := int(digits[j] - '0')
			numbers[i][counter] = intDigit
			counter--

		}
	}

	// DEBUG
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Printf("numbers[%v]: %v\n", i, numbers[i])
	// }

	// Step 3 : Rotate array, starting from the right each column becomes a new row
	rotated := [][]int{}
	for j := 3; j >= 0; j-- { // Start from the end

		newRow := []int{}
		for i := 0; i < len(numbers); i++ {
			if numbers[i][j] != 0 {
				newRow = append(newRow, numbers[i][j])
			}
		}

		rotated = append(rotated, newRow)
	}

	// for i := 0; i < len(rotated); i++ {
	// 	fmt.Printf("rotated[%v]: %v\n", i, rotated[i])
	// }

	total := 0
	if operation == "*" {
		total = 1
	}

	// Step 4 : Create the new integers and do the operation
	for i := 0; i < len(rotated); i++ {

		newNumber := ""
		for j := 0; j < len(rotated[i]); j++ {
			newNumber += strconv.Itoa(rotated[i][j])
		}

		intValue, _ := strconv.Atoi(newNumber)
		if intValue == 0 {
			continue
		}

		fmt.Println(rotated[i], intValue, operation)
		if operation == "*" {
			total *= intValue
		} else {
			total += intValue
		}
	}

	fmt.Printf("total: %v\n", total)
	return total

}

/*

Note to future self, me probably tomororw

You have lots of 2d arrays (each column has a 2d array).
You basically need a method that gets this 2d array, along with its operation
Then you need to rotate the 2d array counter clockwise 90'
and then you can simply do the math. You don't need a new 3d array to do this, I think it's better to loop through
the columns, and treat each column as a separate 2d array. godspeed

*/

func part22(data []string) int {

	// Step 1 : Parse numbers into a 2d array, and operations to a 1D array.
	numberOfRows := len(data) - 1
	numberOfCols := len(strings.Fields(data[0]))

	numbers := make([][]string, numberOfRows)
	operations := make([]string, numberOfCols)

	for i := 0; i <= numberOfRows; i++ {

		rowArray := strings.Fields(data[i])

		// Last row
		if i == numberOfRows {
			copy(operations, rowArray)
			break
		}

		// Rest of the rows should create a 2D array
		fmt.Println(rowArray)
		numbers[i] = rowArray
	}

	fmt.Println(numbers)
	fmt.Println(operations)

	// Create a 3D array to hold the new numbers, init with 0
	copyNumbers := make([][][]int, numberOfRows)
	for i := 0; i < numberOfRows; i++ {

		copyNumbers[i] = make([][]int, numberOfCols)
		for j := 0; j < numberOfCols; j++ {
			y := make([]int, 4) // Max number ?
			copyNumbers[i][j] = y
		}
	}

	// Loop through the actual numbers and try and populate the copyNumbers array
	for j := 0; j < numberOfCols; j++ {

		for i := 0; i < numberOfRows; i++ {

			// Convert string to rune array
			digits := []rune(numbers[i][j])

			// Loop backwards for each element
			counter := 0
			for k := len(digits) - 1; k >= 0; k-- {

				integer := int(digits[k] - '0')

				copyNumbers[i+counter][j][k] = integer
				counter++
			}
		}

		for i := 0; i < numberOfRows; i++ {
			fmt.Println(copyNumbers[i])
		}

		os.Exit(1)
	}

	return 0

}

func part2(data []string) int {

	total := 0

	rows := len(data) - 1 // remove one since the last one contains the operations
	columns := len(strings.Fields(data[0]))

	numbers := make([][]int, rows)
	operations := make([]string, columns)

	for i, row := range data {

		// Remove the whitespace and split
		row := strings.Fields(row)

		if i == rows {
			for j, v := range row {
				operations[j] = v
			}
			break
		}

		// Create local row numbers
		rowNumbers := make([]int, columns)
		for j, v := range row {

			intNumber, _ := strconv.Atoi(v)

			rowNumbers[j] = intNumber
		}

		numbers[i] = rowNumbers

	}

	for j := 0; j < columns-1; j++ {

		println("new line")

		for i := 0; i < rows; i++ {

			println(numbers[i][j])

		}

	}

	return total
}

func main() {

	data := parseFile("06.txt")
	// fmt.Println(part1(data))
	// fmt.Println(part2(data))
	fmt.Println(part222(data))
}
