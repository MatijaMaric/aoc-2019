package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

func main() {
	code := utils.ReadIntCode("input.txt")
	grid, dim := part1(code)
	part2(code, grid, dim)
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
			dim = vec2{X: x - 1, Y: y}
		}
	}

	part1 := 0

	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			pos := vec2{X: x, Y: y}
			if grid[pos] == '#' {
				if countAdjacent(grid, pos) == 4 {
					part1 += x * y
				}
			}
		}
	}

	fmt.Println(part1)
	return
}

func part2(code []int, grid map[vec2]rune, dim vec2) {
	code[0] = 2
	input, output := make(chan int, 1), make(chan int)

	go utils.IntCodeMachine(code, input, output)

	instructions := getInstructions(grid, dim)
	fmt.Println(instructions)
}

func getRobot(grid map[vec2]rune) (pos, orientation vec2) {
	for k, v := range grid {
		switch v {
		case '^':
			return k, vec2{Y: -1}
		case 'v':
			return k, vec2{Y: 1}
		case '>':
			return k, vec2{X: 1}
		case '<':
			return k, vec2{X: -1}
		}
	}
	panic("nooooo")
}

func getRotation(current, next vec2) rune {
	if current.X == 1 {
		if next.Y == 1 {
			return 'R'
		} else {
			return 'L'
		}
	} else {
		if next.Y == 1 {
			return 'L'
		} else {
			return 'R'
		}
	}
}

func getInstructions(grid map[vec2]rune, dim vec2) []string {
	var ans []string
	pos, orientation := getRobot(grid)
	grid[pos] = '.'

	moves := 0
	var instruction rune

	printGrid(grid, dim)
	fmt.Println()

	for {
		next := pos.Add(orientation)
		if grid[next] == '#' {
			if countAdjacent(grid, next) < 3 {
				grid[next] = '.'
			}
			moves++
			pos = next
		} else {
			printGrid(grid, dim)
			fmt.Println()
			new := fmt.Sprintf("%c,%d,", instruction, moves)
			ans = append(ans, new)
			if grid[pos.Add(vec2{X: 1})] == '#' {
				instruction = getRotation(orientation, vec2{X: 1})
				orientation = vec2{X: 1}
			} else if grid[pos.Add(vec2{X: -1})] == '#' {
				instruction = getRotation(orientation, vec2{X: -1})
				orientation = vec2{X: -1}
			} else if grid[pos.Add(vec2{Y: 1})] == '#' {
				instruction = getRotation(orientation, vec2{Y: 1})
				orientation = vec2{Y: 1}
			} else if grid[pos.Add(vec2{Y: -1})] == '#' {
				instruction = getRotation(orientation, vec2{Y: -1})
				orientation = vec2{Y: -1}
			} else {
				break
			}
		}
	}

	return ans[1:]
}

func printGrid(grid map[vec2]rune, dim vec2) {
	for y := 0; y <= dim.Y; y++ {
		var builder strings.Builder
		for x := 0; x <= dim.X; x++ {
			builder.WriteRune(grid[vec2{X: x, Y: y}])
		}
		fmt.Println(builder.String())
	}
}

func countAdjacent(grid map[vec2]rune, pos vec2) (ans int) {
	if grid[pos.Add(vec2{X: 1})] == '#' {
		ans++
	}
	if grid[pos.Add(vec2{X: -1})] == '#' {
		ans++
	}
	if grid[pos.Add(vec2{Y: 1})] == '#' {
		ans++
	}
	if grid[pos.Add(vec2{Y: -1})] == '#' {
		ans++
	}
	return
}
