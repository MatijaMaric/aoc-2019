package utils

// ReadIntCode reads IntCode from file
func ReadIntCode(path string) []int {
	return IntList(ReadLines(path)[0])
}

func pow10(x int) int {
	ans := 1
	for i := 0; i < x; i++ {
		ans *= 10
	}
	return ans
}

func getMode(instruction int, pos int) int {
	return instruction % pow10(pos+2) / pow10(pos+1)
}

func getArg(memory map[int]int, pc int, pos int, rel int, read bool) int {
	instruction := memory[pc]
	mode := getMode(instruction, pos)
	arg := memory[pc+pos]
	if mode == 0 && read {
		arg = memory[arg]
	} else if mode == 2 && read {
		arg = memory[arg+rel]
	} else if mode == 2 && !read {
		arg += rel
	}
	return arg
}

func unaryInstruction(memory map[int]int, pc int, rel int) (arg, res int) {
	arg = getArg(memory, pc, 1, rel, true)
	res = getArg(memory, pc, 2, rel, false)
	return
}

func binaryInstruction(memory map[int]int, pc int, rel int) (arg1, arg2, res int) {
	arg1 = getArg(memory, pc, 1, rel, true)
	arg2 = getArg(memory, pc, 2, rel, true)
	res = getArg(memory, pc, 3, rel, false)
	return
}

// IntCodeMachine emulates IntCode interpreter
func IntCodeMachine(program []int, input chan int, output chan int) {
	memory := ArrayToMap(program)

	rel := 0
	pc := 0
	for {
		instruction := memory[pc]
		opcode := instruction % 100
		switch opcode {
		case 1:
			a, b, res := binaryInstruction(memory, pc, rel)
			memory[res] = a + b
			pc += 4
		case 2:
			a, b, res := binaryInstruction(memory, pc, rel)
			memory[res] = a * b
			pc += 4
		case 3:
			addr := getArg(memory, pc, 1, rel, false)
			memory[addr] = <-input
			pc += 2
		case 4:
			a, _ := unaryInstruction(memory, pc, rel)
			output <- a
			pc += 2
		case 5:
			a, b, _ := binaryInstruction(memory, pc, rel)
			if a != 0 {
				pc = b
			} else {
				pc += 3
			}
		case 6:
			a, b, _ := binaryInstruction(memory, pc, rel)
			if a == 0 {
				pc = b
			} else {
				pc += 3
			}
		case 7:
			a, b, res := binaryInstruction(memory, pc, rel)
			if a < b {
				memory[res] = 1
			} else {
				memory[res] = 0
			}
			pc += 4
		case 8:
			a, b, res := binaryInstruction(memory, pc, rel)
			if a == b {
				memory[res] = 1
			} else {
				memory[res] = 0
			}
			pc += 4
		case 9:
			a, _ := unaryInstruction(memory, pc, rel)
			rel += a
			pc += 2
		case 99:
			close(output)
			return
		default:
			panic(opcode)
		}
	}
}
