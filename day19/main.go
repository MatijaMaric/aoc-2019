package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type drone struct {
	utils.IntCodeVM
}

func main() {
	// var vm drone
	// vm.InitFromFile("input.txt")

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
	var vm utils.IntCodeVM
	vm.Init(code)
	vm.Write(x)
	vm.Write(y)

	return vm.Read() == 1
}
