package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	code := utils.IntList(utils.ReadLines("input.txt")[0])
	input, output := make(chan int, 1), make(chan int)

	input <- 2

	go utils.IntCodeMachine(code, input, output)

	for out := range output {
		fmt.Println(out)
	}
}
