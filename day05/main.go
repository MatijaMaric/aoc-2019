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

func parseInstruction(memory []int, pc int) (arg1, arg2, target int) {
	instruction := memory[pc]
	arg1 = memory[pc+1]
	if (instruction%1000)/100 == 0 {
		arg1 = memory[arg1]
	}
	arg2 = memory[pc+2]
	if (instruction%10000)/1000 == 0 {
		arg2 = memory[arg2]
	}
	target = memory[pc+3]
	return
}

func intCodeMachine(program []int, input int) int {
	memory := make([]int, len(program))
	copy(memory, program)

	pc := 0
	output := 0
	for {
		instruction := memory[pc]
		opcode := instruction % 100
		switch opcode {
		case 1:
			a, b, target := parseInstruction(memory, pc)
			memory[target] = a + b
			pc += 4
		case 2:
			a, b, target := parseInstruction(memory, pc)
			memory[target] = a * b
			pc += 4
		case 3:
			memory[memory[pc+1]] = input
			pc += 2
		case 4:
			arg := memory[pc+1]
			if instruction%1000/100 == 0 {
				arg = memory[arg]
			}
			output = arg
			pc += 2
		case 99:
			return output
		default:
			panic("nooo")
		}
	}
}
