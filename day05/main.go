package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	codes := utils.IntList(utils.ReadLines("input.txt")[0])
	output := intCodeMachine(codes, 1)
	fmt.Println(output)
}

func flags(instruction int) (b bool, a bool) {
	number := instruction / 100
	b = number%10 == 1
	number = number / 10
	a = number%10 == 1
	return
}

func intCodeMachine(program []int, input int) int {
	pc := 0
	output := 0
	for program[pc]%100 != 99 {
		opcode := program[pc]
		b, a := flags(opcode)
		opcode = program[pc] % 100
		x := program[program[pc+1]]
		if a {
			x = program[pc+1]
		}
		y := program[program[pc+2]]
		if b {
			y = program[pc+2]
		}
		switch opcode {
		case 1:
			program[program[pc+3]] = x + y
			pc += 4
		case 2:
			program[program[pc+3]] = x * y
			pc += 4
		case 3:
			program[program[pc+1]] = input
			pc += 2
		case 4:
			output = program[program[pc+1]]
			pc += 2
		}
	}
	return output
}
