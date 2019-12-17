package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

func main() {
	code := utils.ReadIntCode("input.txt")
	part1(code)
}

func part1(code []int) (grid map[vec2]rune, dim vec2) {
	input, output := make(chan int, 1), make(chan int)

	go utils.IntCodeMachine(code, input, output)

	grid = make(map[vec2]rune)

	var x, y int
	for char := range output {
		if char == '\n' {
			y++
			x = 0
		} else {
			grid[vec2{X: x, Y: y}] = rune(char)
			x++
		}
		if char != '\n' {
			dim = vec2{X: x, Y: y}
		}
	}

	part1 := 0

	up := vec2{Y: -1}
	down := vec2{Y: 1}
	left := vec2{X: -1}
	right := vec2{X: 1}
	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			pos := vec2{X: x, Y: y}
			if grid[pos] == '#' {
				if grid[pos.Add(up)] == '#' && grid[pos.Add(down)] == '#' && grid[pos.Add(left)] == '#' && grid[pos.Add(right)] == '#' {
					part1 += x * y
				}
			}
		}
	}

	fmt.Println(part1)
	return
}
