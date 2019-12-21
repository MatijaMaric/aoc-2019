package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	code := utils.ReadIntCode("input.txt")

	fmt.Println(part1(code))
	fmt.Println(part2(code))
}

func part1(code []int) int {
	count := 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			if ping(code, x, y) {
				count++
			}
		}
	}
	return count
}

func part2(code []int) int {
	x, y := 0, 99
	for {
		fmt.Println(x, y)
		if ping(code, x, y) {
			if ping(code, x+99, y-99) {
				return 10000*x + y - 99
			}
			x--
			y++
		}
		x++
	}
}

func ping(code []int, x, y int) bool {
	input, output := make(chan int, 1), make(chan int)

	go utils.IntCodeMachine(code, input, output)

	input <- x
	input <- y

	return <-output == 1
}
