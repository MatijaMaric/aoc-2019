package main

func generatePermutations(array []int) [][]int {
	var permute func([]int, int)
	var ans [][]int

	permute = func(array []int, n int) {
		if n == 1 {
			tmp := make([]int, len(array))
			copy(tmp, array)
			ans = append(ans, tmp)
		} else {
			for i := 0; i < n; i++ {
				permute(array, n-1)
				if n%2 == 1 {
					array[i], array[n-1] = array[n-1], array[i]
				} else {
					array[0], array[n-1] = array[n-1], array[0]
				}
			}
		}
	}

	permute(array, len(array))
	return ans
}

func runAmplifiers(code []int, sequence []int) int {
	a := make(chan int, 1)
	b := make(chan int)
	c := make(chan int)
	d := make(chan int)
	e := make(chan int)

	halt := make(chan bool)

	go intCodeMachine(code, a, b, halt)
	go intCodeMachine(code, b, c, halt)
	go intCodeMachine(code, c, d, halt)
	go intCodeMachine(code, d, e, halt)
	go intCodeMachine(code, e, a, halt)

	a <- sequence[0]
	b <- sequence[1]
	c <- sequence[2]
	d <- sequence[3]
	e <- sequence[4]

	a <- 0

	for i := 0; i < 5; i++ {
		<-halt
	}

	return <-a
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

func intCodeMachine(program []int, input chan int, output chan int, halt chan bool) {
	memory := make([]int, len(program))
	copy(memory, program)

	pc := 0
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
			memory[memory[pc+1]] = <-input
			pc += 2
		case 4:
			arg := memory[pc+1]
			if instruction%1000/100 == 0 {
				arg = memory[arg]
			}
			output <- arg
			pc += 2
		case 5:
			a, b, _ := parseInstruction(memory, pc)
			if a != 0 {
				pc = b
			} else {
				pc += 3
			}
		case 6:
			a, b, _ := parseInstruction(memory, pc)
			if a == 0 {
				pc = b
			} else {
				pc += 3
			}
		case 7:
			a, b, target := parseInstruction(memory, pc)
			if a < b {
				memory[target] = 1
			} else {
				memory[target] = 0
			}
			pc += 4
		case 8:
			a, b, target := parseInstruction(memory, pc)
			if a == b {
				memory[target] = 1
			} else {
				memory[target] = 0
			}
			pc += 4
		case 99:
			halt <- true
			return
		default:
			panic("nooo")
		}
	}
}
