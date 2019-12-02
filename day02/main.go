package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

const target = 19690720

func main() {
	lines := utils.ReadLines("input.txt")
	resetCode := utils.IntList(lines[0])
	numCode := make([]int, len(resetCode))
	{
		copy(numCode, resetCode)

		numCode[1] = 12
		numCode[2] = 2

		intCodeMachine(numCode)
		fmt.Println(numCode[0])
	}
	{
		for i := 0; i <= 99; i++ {
			for j := 0; j <= 99; j++ {
				copy(numCode, resetCode)

				numCode[1] = i
				numCode[2] = j

				intCodeMachine(numCode)
				if numCode[0] == target {
					fmt.Println(100*i + j)
				}
			}
		}
	}

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
