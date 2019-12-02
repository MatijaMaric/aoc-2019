package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

const target = 19690720

func main() {
	lines := utils.ReadLines("input.txt")
	code := utils.IntList(lines[0])

	ans := runMachine(code, 12, 2)
	fmt.Println(ans)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			ans := runMachine(code, i, j)
			if ans == target {
				fmt.Println(100*i + j)
			}
		}
	}

}

func runMachine(input []int, noun int, verb int) int {
	code := make([]int, len(input))
	copy(code, input)
	code[1] = noun
	code[2] = verb
	intCodeMachine(code)
	return code[0]
}

func intCodeMachine(input []int) {
	pc := 0
	for input[pc] != 99 {
		opcode := input[pc]
		switch opcode {
		case 1:
			a := input[input[pc+1]]
			b := input[input[pc+2]]
			input[input[pc+3]] = a + b
		case 2:
			a := input[input[pc+1]]
			b := input[input[pc+2]]
			input[input[pc+3]] = a * b
		}
		pc += 4
	}
}
