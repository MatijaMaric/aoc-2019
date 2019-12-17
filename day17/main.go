package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

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
	compressed := compress(strings.Join(instructions, ""))
	// fmt.Println(compressed)

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

// pobogu zašto
func getRotation(current, next vec2) rune {
	if current.Eq(vec2{Y: -1}) {
		if next.Eq(vec2{X: -1}) {
			return 'L'
		}
		if next.Eq(vec2{X: 1}) {
			return 'R'
		}
	}
	if current.Eq(vec2{X: 1}) {
		if next.Eq(vec2{Y: -1}) {
			return 'L'
		}
		if next.Eq(vec2{Y: 1}) {
			return 'R'
		}
	}
	if current.Eq(vec2{Y: 1}) {
		if next.Eq(vec2{X: 1}) {
			return 'L'
		}
		if next.Eq(vec2{X: -1}) {
			return 'R'
		}
	}
	if current.Eq(vec2{X: -1}) {
		if next.Eq(vec2{Y: 1}) {
			return 'L'
		}
		if next.Eq(vec2{Y: -1}) {
			return 'R'
		}
	}
	panic("nooo")
}

func getInstructions(grid map[vec2]rune, dim vec2) []string {
	var ans []string
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
			new := fmt.Sprintf("%c,%d,", instruction, moves)
			ans = append(ans, new)
			moves = 0
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
