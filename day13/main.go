package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

func main() {
	code := utils.ReadIntCode("input.txt")

	blocks := countBlocks(code)
	score := playGame(code)

	fmt.Println(blocks)
	fmt.Println(score)
}

func countBlocks(code []int) int {
	input, output := make(chan int), make(chan int)
	go utils.IntCodeMachine(code, input, output)

	blocks := 0

	for {
		_, ok := <-output
		if !ok {
			break
		}
		_, id := <-output, <-output

		if id == 2 {
			blocks++
		}
	}
	return blocks
}

func playGame(code []int) int {
	input, output := make(chan int, 1), make(chan int)
	code[0] = 2
	go utils.IntCodeMachine(code, input, output)

	grid := make(map[vec2]int)

	var ball, paddle int

	for {
		x, ok := <-output
		if !ok {
			break
		}
		y, id := <-output, <-output

		if x == -1 {
			return id
		}

		pos := vec2{X: x, Y: y}
		grid[pos] = id

		switch id {
		case 3:
			paddle = x
		case 4:
			ball = x
			if ball < paddle {
				input <- -1
			} else if ball > paddle {
				input <- 1
			} else {
				input <- 0
			}
		}
	}
	return -1
}

// func printGame(grid map[vec2]rune, min, max vec2) {
// 	for y := min.Y; y <= max.Y; y++ {
// 		var builder strings.Builder
// 		for x := min.X; x <= max.X; x++ {
// 			builder.WriteRune(grid[vec2{X: x, Y: y}])
// 		}
// 		fmt.Println(builder.String())
// 	}
// }
