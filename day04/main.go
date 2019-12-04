package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	input := utils.ReadLines("input.txt")[0]
	from, to := parseInput(input)
	fmt.Println(solve(from, to))
}

func parseInput(input string) (int, int) {
	str := strings.Split(input, "-")
	return utils.ToInt(str[0]), utils.ToInt(str[1])
}

func solve(from, to int) int {
	cnt := 0
	for i := from; i <= to; i++ {
		if isValid(i) {
			cnt++
		}
	}
	return cnt
}

func isValid(number int) bool {
	str := utils.IntToStr(number)
	hasDoubles := false
	increases := true
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			hasDoubles = true
		}
		if str[i] > str[i+1] {
			increases = false
		}
	}
	return hasDoubles && increases
}
