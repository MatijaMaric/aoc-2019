package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

const (
	north = iota + 1
	south
	west
	east
)

const (
	wall = iota
	moved
	oxygen
)

type mazeState struct {
	pos  vec2
	path []int
}

func main() {
	code := utils.ReadIntCode("input.txt")
	dist, grid, oxygen := part1(code)
	fmt.Println(dist)
	fillTime := part2(grid, oxygen)
	fmt.Println(fillTime)
}

func part1(code []int) (oxygenPath int, grid map[vec2]rune, oxygenPos vec2) {
	queue := []mazeState{mazeState{vec2{}, []int{}}}
	visited := make(map[vec2]bool)
	grid = make(map[vec2]rune)
	grid[queue[0].pos] = '.'
	visited[queue[0].pos] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for move := 1; move <= 4; move++ {
			next := current.pos.Add(move2vec(move))
			if visited[next] {
				continue
			}
			visited[next] = true
			nextPath := make([]int, len(current.path))
			copy(nextPath, current.path)
			nextPath = append(nextPath, move)
			status := runRobot(code, nextPath)
			switch status {
			case wall:
				grid[next] = '#'
			case moved:
				grid[next] = '.'
				queue = append(queue, mazeState{next, nextPath})
			case oxygen:
				grid[next] = 'O'
				queue = append(queue, mazeState{next, nextPath})
				oxygenPath = len(nextPath)
				oxygenPos = next
			default:
				panic("nooo")
			}
		}
	}
	return
}

func part2(grid map[vec2]rune, oxygenPos vec2) int {
	queue := []vec2{oxygenPos}
	visited := make(map[vec2]int)
	visited[queue[0]] = 0
	ans := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for move := 1; move <= 4; move++ {
			next := current.Add(move2vec(move))

			if _, ok := visited[next]; ok {
				continue
			} else {
				ans = utils.Max(ans, visited[current]+1)
				if grid[next] == '.' {
					queue = append(queue, next)
				}
			}
			visited[next] = visited[current] + 1
		}
	}
	return ans
}

func runRobot(code []int, path []int) (status int) {
	input, output := make(chan int, 1), make(chan int)
	go utils.IntCodeMachine(code, input, output)
	for i, move := range path {
		input <- move
		status = <-output
		if status != moved && i < len(path)-1 {
			fmt.Println("YOOOO")
		}
	}
	return
}

func move2vec(move int) vec2 {
	switch move {
	case north:
		return vec2{X: 0, Y: -1}
	case east:
		return vec2{X: 1, Y: 0}
	case south:
		return vec2{X: 0, Y: 1}
	case west:
		return vec2{X: -1, Y: 0}
	}
	panic("nooo")
}
