package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

func main() {
	code := utils.IntList(utils.ReadLines("input.txt")[0])

	canvas := make(map[vec2]int)
	run(code, canvas, 0)

	fmt.Println(len(canvas))

	canvas = make(map[vec2]int)
	min, max := run(code, canvas, 1)

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if canvas[vec2{X: x, Y: y}] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func run(code []int, canvas map[vec2]int, in int) (min, max vec2) {
	input, output := make(chan int, 1), make(chan int)

	input <- in

	var pos vec2
	direction := vec2{X: 0, Y: -1}

	go utils.IntCodeMachine(code, input, output)

	for {
		color, ok := <-output
		if !ok {
			break
		}
		rotate := <-output
		canvas[pos] = color
		if rotate == 0 {
			direction = rotateLeft(direction)
		} else {
			direction = rotateRight(direction)
		}
		pos = pos.Add(direction)
		min = minVec(min, pos)
		max = maxVec(max, pos)

		input <- canvas[pos]
	}
	return
}

func rotateLeft(v vec2) vec2 {
	if v.X == 1 {
		return vec2{X: 0, Y: -1}
	}
	if v.Y == -1 {
		return vec2{X: -1, Y: 0}
	}
	if v.X == -1 {
		return vec2{X: 0, Y: 1}
	}
	if v.Y == 1 {
		return vec2{X: 1, Y: 0}
	}
	panic("invalid move")
}

func rotateRight(v vec2) vec2 {
	return rotateLeft(v).Multiply(-1)
}

func minVec(a, b vec2) vec2 {
	return vec2{X: utils.Min(a.X, b.X), Y: utils.Min(a.Y, b.Y)}
}

func maxVec(a, b vec2) vec2 {
	return vec2{X: utils.Max(a.X, b.X), Y: utils.Max(a.Y, b.Y)}
}
