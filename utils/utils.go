package utils

import (
	"bufio"
	"os"
	"strconv"
)

// Check panics if error is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadLines reads a file into an string array
func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())
	return lines
}

// ReadNumbers reads numbers from file to int array
func ReadNumbers(filename string) []int {
	lines := ReadLines(filename)
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = ToInt(line)
	}
	return numbers
}

// ToInt converts string to integer
func ToInt(text string) int {
	x, err := strconv.Atoi(text)
	Check(err)
	return x
}
