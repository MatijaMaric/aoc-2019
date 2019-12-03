package main

import (
	"fmt"
	"strings"
	"math"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type Point struct {
	x	int
	y	int
}

type Segment struct {
	direction	Point
	length		int
}

type Wire []Segment

func main() {
	paths := utils.ReadLines("input.txt")
	wires := make([]Wire, len(paths))
	for i, path := range paths {
		wires[i] = parseWire(path)
	}
	
	grid := make(map[Point]int)
	var pos Point
	
	var steps int
	
	for _, segment := range wires[0] {
		for i := 0; i < segment.length; i++ {
			pos.x += segment.direction.x
			pos.y += segment.direction.y
			steps++
			grid[pos] = steps
		}
	}
	
	minManhattan := math.MaxInt32
	minSteps := math.MaxInt32
	pos = Point{0, 0}
	steps = 0
	for _, segment := range wires[1] {
		for i := 0; i < segment.length; i++ {
			pos.x += segment.direction.x
			pos.y += segment.direction.y
			steps++
			if grid[pos] > 0 {
				distance := manhattan(pos)
				minManhattan = utils.Min(minManhattan, distance)
				minSteps = utils.Min(minSteps, grid[pos] + steps)
			}
		}
	}

	fmt.Println(minManhattan)
	fmt.Println(minSteps)

}

func parseWire(path string) Wire {
	var lines []Segment
	for _, line := range strings.Split(path, ",") {
		var direction Point
		length := utils.ToInt(line[1:])
		switch line[0] {
		case 'U':
			direction.y--
		case 'D':
			direction.y++
		case 'L':
			direction.x--
		case 'R':
			direction.x++
		}
		lines = append(lines, Segment{direction, length})
	}
	return lines
}

func manhattan(p Point) int {
	return utils.Abs(p.x) + utils.Abs(p.y)
}