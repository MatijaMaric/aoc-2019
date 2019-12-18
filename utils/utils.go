package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// IntToStr converts integer to string
func IntToStr(number int) string {
	return strconv.Itoa(number)
}

// IntList converts comma separated string into integer array
func IntList(list string) []int {
	var numbers []int
	for _, num := range strings.Split(list, ",") {
		numbers = append(numbers, ToInt(num))
	}
	return numbers
}

// Abs gets absolute value of integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min gets minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max gets maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ArrayToMap converts an integer array to map
func ArrayToMap(array []int) map[int]int {
	ans := make(map[int]int)
	for i, val := range array {
		ans[i] = val
	}
	return ans
}

// Gcd returns greatest common divisor
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Sgn return sign of number
func Sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

// Lcm returns least common multiple of two numbers
func Lcm(a, b int) int {
	return Abs(a*b) / Gcd(a, b)
}

// GridToMap converts array of strings to map (by coords) of characters
func GridToMap(grid []string) map[Vector2D]rune {
	ans := make(map[Vector2D]rune)
	for y, row := range grid {
		for x, v := range row {
			ans[Vector2D{x, y}] = v
		}
	}
	return ans
}

// ReadGrid reads array of strings from file into a map (by coords) of characters
func ReadGrid(path string) map[Vector2D]rune {
	return GridToMap(ReadLines(path))
}

// GridSize returns dimensions of grid (max x,y coords)
func GridSize(grid map[Vector2D]rune) Vector2D {
	var ans Vector2D
	for k := range grid {
		ans = MaxVec2(ans, k)
	}
	return ans.Add(Vector2D{1, 1})
}

// PrintGrid prints grid to stdout
func PrintGrid(grid map[Vector2D]rune) {
	dims := GridSize(grid)
	for y := 0; y < dims.Y; y++ {
		for x := 0; x < dims.X; x++ {
			fmt.Printf("%c", grid[Vector2D{x, y}])
		}
		fmt.Println()
	}
}

// StrToRunes converts string to array of characters
func StrToRunes(str string) []rune {
	ans := make([]rune, len(str))
	for i, c := range str {
		ans[i] = c
	}
	return ans
}

// RunesToStr converts array of characters to string
func RunesToStr(runes []rune) string {
	var builder strings.Builder
	for _, c := range runes {
		builder.WriteRune(c)
	}
	return builder.String()
}

// AppendRune appends a character to end of string
func AppendRune(str string, c rune) string {
	var builder strings.Builder
	builder.WriteString(str)
	builder.WriteRune(c)
	return builder.String()
}
