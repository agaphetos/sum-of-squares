package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSumOfSquares(input []string, iterator, currentSum int) int {
	// recursion terminator condition
	if iterator >= len(input) {
		return currentSum
	}

	number, err := strconv.Atoi(input[iterator])
	// validate number conversion and if positive integer add squared number
	if err == nil && number >= 0 && number <= 100 {
		currentSum += number * number
	}

	return getSumOfSquares(input, iterator+1, currentSum)
}

func exec(scanner *bufio.Scanner, testCases, iterator int, results []int) []int {
	// recursion terminator condition
	if iterator < testCases {
		// get test case input
		scanner.Scan()
		numberCount, err := strconv.Atoi(scanner.Text())
		// validate input number count
		if err == nil && numberCount > 0 && numberCount <= 100 {
			// get input numbers
			scanner.Scan()
			input := scanner.Text()
			inputFields := strings.Fields(input)
			results[iterator] = getSumOfSquares(inputFields, 0, 0)
		} else {
			fmt.Println("Invalid Input")
		}
		exec(scanner, testCases, iterator+1, results)
	}

	return results
}

func printResults(results []int, iterator int) {
	if iterator >= len(results) {
		return
	}
	fmt.Println(results[iterator])
	printResults(results, iterator+1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	testCases, err := strconv.Atoi(scanner.Text())
	if err == nil && testCases >= 0 && testCases <= 100 {
		var results = make([]int, testCases)
		sumOfSquares := exec(scanner, testCases, 0, results)
		printResults(sumOfSquares, 0)
	} else {
		fmt.Println("Invalid Input")
	}
}
