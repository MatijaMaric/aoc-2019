package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

var (
	up    = vec2{Y: -1}
	right = vec2{X: 1}
	down  = vec2{Y: 1}
	left  = vec2{X: -1}
)

var directions = []vec2{up, right, down, left}

type vec2 = utils.Vector2D

type node struct {
	pos    vec2
	to     []vec2
	offset int
}

func main() {
	grid := utils.ReadGrid("test.txt")
	dim := utils.GridSize(grid)
	nodes, start, end := getNodes(grid)
	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			if v, ok := nodes[vec2{X: x, Y: y}]; ok && v.offset != 0 {
				if v.offset > 0 {
					fmt.Print("^")
				} else {
					fmt.Print("v")
				}
			} else {
				fmt.Printf("%c", grid[vec2{X: x, Y: y}])
			}
		}
		fmt.Println()
	}
	fmt.Println(bfs(nodes, start, end, false, grid))
	fmt.Println(bfs(nodes, start, end, true, grid))
}

type position struct {
	vec   vec2
	level int
}

func bfs(nodes map[vec2]*node, start, end vec2, recursive bool, grid map[vec2]rune) int {
	type item struct {
		pos  position
		dist int
		path []position
	}
	q := []item{item{position{start, 0}, 0, []position{position{start, 0}}}}
	visited := make(map[position]bool)
	visited[position{start, 0}] = true

	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		printGrid(grid, current.path)
		if current.pos.vec.Eq(end) && current.pos.level == 0 {
			return current.dist
		}

		// fmt.Println(current)
		for _, next := range nodes[current.pos.vec].to {
			level := current.pos.level
			if recursive {
				level += nodes[current.pos.vec].offset
			}
			nextPos := position{next, level}
			dist := current.dist + 1
			if level >= 0 && !visited[nextPos] {
				visited[nextPos] = true
				q = append(q, item{nextPos, dist, append(current.path, nextPos)})
			}
		}
	}

	panic("noooo")
}

func getNodes(grid map[vec2]rune) (nodes map[vec2]*node, start, end vec2) {
	nodes = make(map[vec2]*node)
	portals := make(map[string]([]vec2))
	for pos, c := range grid {
		if c == '.' {
			nodes[pos] = &node{pos, []vec2{}, 0}
			for _, dir := range directions {
				neighbor := pos.Add(dir)
				if grid[neighbor] == '.' {
					nodes[pos].to = append(nodes[pos].to, neighbor)
				}
			}
		}
		if n, p := grid[pos.Add(up)], grid[pos.Add(down)]; isLetter(c) && (isLetter(n) || isLetter(p)) {
			if n == '.' {
				label := fmt.Sprintf("%c%c", c, p)
				portals[label] = append(portals[label], pos.Add(up))
			}
			if p == '.' {
				label := fmt.Sprintf("%c%c", n, c)
				portals[label] = append(portals[label], pos.Add(down))
			}
		}
		if n, p := grid[pos.Add(left)], grid[pos.Add(right)]; isLetter(c) && (isLetter(n) || isLetter(p)) {
			if n == '.' {
				label := fmt.Sprintf("%c%c", c, p)
				portals[label] = append(portals[label], pos.Add(left))
			}
			if p == '.' {
				label := fmt.Sprintf("%c%c", n, c)
				portals[label] = append(portals[label], pos.Add(right))
			}
		}
	}
	// fmt.Println(portals)
	dim := utils.GridSize(grid)
	for label, pos := range portals {
		if strings.Compare(label, "AA") == 0 {
			start = pos[0]
			continue
		} else if strings.Compare(label, "ZZ") == 0 {
			end = pos[0]
			continue
		}
		if pos[0].X == 2 || pos[0].Y == 2 || pos[0].X == dim.X-3 || pos[0].Y == dim.Y-3 {
			nodes[pos[0]].offset = -1
		} else {
			nodes[pos[0]].offset = 1
		}
		nodes[pos[0]].to = append(nodes[pos[0]].to, pos[1])
		if pos[1].X == 2 || pos[1].Y == 2 || pos[1].X == dim.X-3 || pos[1].Y == dim.Y-3 {
			nodes[pos[1]].offset = -1
		} else {
			nodes[pos[1]].offset = 1
		}
		nodes[pos[1]].to = append(nodes[pos[1]].to, pos[0])
	}
	return
}

func isLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func printGrid(grid map[vec2]rune, path []position) {
	visited := make(map[vec2]rune)
	for _, v := range path {
		visited[v.vec] = rune(v.level) + '1'
	}
	dim := utils.GridSize(grid)
	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			if v, ok := visited[vec2{X: x, Y: y}]; ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf("%c", grid[vec2{X: x, Y: y}])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
