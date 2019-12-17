package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

var (
	up    = vec2{Y: -1}
	right = vec2{X: 1}
	down  = vec2{Y: 1}
	left  = vec2{X: -1}
)

var directions []vec2 = []vec2{up, right, down, left}

func main() {
	code := utils.ReadIntCode("input.txt")
	grid, dim := part1(code)
	dust := part2(code, grid, dim)
	fmt.Println(dust)
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

	// printGrid(grid, dim)

	fmt.Println(part1)
	return
}

func part2(code []int, grid map[vec2]rune, dim vec2) int {
	code[0] = 2
	input, output := make(chan int, 1000), make(chan int)

	go utils.IntCodeMachine(code, input, output)

	instructions := getInstructions(grid, dim)
	// fmt.Println(strings.Join(instructions, ""))
	compressed := compress(instructions)
	fmt.Println(compressed)

	for _, x := range compressed {
		input <- int(x)
	}
	input <- 'n'
	input <- '\n'

	for dust := range output {
		if dust < 256 {
			fmt.Printf("%c", rune(dust))
		} else {
			return dust
		}
	}
	return -1
}

func getRobot(grid map[vec2]rune) (pos, orientation vec2) {
	for k, v := range grid {
		switch v {
		case '^':
			return k, up
		case 'v':
			return k, down
		case '>':
			return k, right
		case '<':
			return k, left
		}
	}
	panic("nooooo")
}

func turnLeft(v vec2) vec2 {
	if v.Eq(up) {
		return left
	}
	if v.Eq(right) {
		return up
	}
	if v.Eq(down) {
		return right
	}
	if v.Eq(left) {
		return down
	}
	panic("nooo")
}

func turnRight(v vec2) vec2 {
	return turnLeft(v).Multiply(-1)
}

func getInstructions(grid map[vec2]rune, dim vec2) string {
	var builder strings.Builder
	pos, orientation := getRobot(grid)
	grid[pos] = '.'

	moves := 0
	var instruction rune

	for {
		next := pos.Add(orientation)
		if grid[next] == '#' {
			if countAdjacent(grid, next) < 3 {
				grid[next] = '.'
			}
			moves++
			pos = next
		} else {
			if moves > 0 {
				builder.WriteString(fmt.Sprintf("%c,%d,", instruction, moves))
			}
			moves = 0
			if turn := turnLeft(orientation); grid[pos.Add(turn)] == '#' {
				instruction = 'L'
				orientation = turn
			} else if turn := turnRight(orientation); grid[pos.Add(turn)] == '#' {
				instruction = 'R'
				orientation = turn
			} else {
				break
			}
		}
	}

	return builder.String()
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
	for _, dir := range directions {
		if grid[pos.Add(dir)] == '#' {
			ans++
		}
	}
	return
}

func compress(original string) string {
	instructionRegex := regexp.MustCompile("[LR],\\d+,")
	instructions := instructionRegex.FindAllString(original, -1)

	patSet := make(map[string]bool)
	for i := 0; i < len(instructions); i++ {
		for j := i; j < len(instructions); j++ {
			comb := strings.Join(instructions[i:j+1], "")
			if len(comb) <= 20 {
				if strings.Count(original, comb) > 1 {
					patSet[comb] = true
				}
			}
		}
	}

	var patterns []string
	for k := range patSet {
		patterns = append(patterns, k)
	}

	for i := 0; i < len(patterns)-2; i++ {
		for j := i + 1; j < len(patterns)-1; j++ {
			for k := j + 1; k < len(patterns); k++ {
				sorted := []string{patterns[i], patterns[j], patterns[k]}
				sort.Slice(sorted, func(a, b int) bool {
					return len(sorted[a]) > len(sorted[b])
				})
				substituted := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(original, sorted[0], "A,"), sorted[1], "B,"), sorted[2], "C,")
				if !strings.ContainsAny(substituted, "LR") {
					var builder strings.Builder

					builder.WriteString(strings.TrimSuffix(substituted, ","))
					builder.WriteRune('\n')
					builder.WriteString(strings.TrimSuffix(sorted[0], ","))
					builder.WriteRune('\n')
					builder.WriteString(strings.TrimSuffix(sorted[1], ","))
					builder.WriteRune('\n')
					builder.WriteString(strings.TrimSuffix(sorted[2], ","))
					builder.WriteRune('\n')

					return builder.String()
				}
			}
		}
	}
	panic("no result")
}
