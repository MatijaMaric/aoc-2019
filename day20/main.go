package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

type portal struct {
	name     string
	from, to vec2
}

func main() {
	grid := utils.ReadGrid("input.txt")
	portals := findPortals(grid)
	edges := findEdges(grid, portals)

	dists := findPath(edges)

	fmt.Println(dists["ZZ"] - 1)
}

func findPath(edges map[string](map[string]int)) map[string]int {
	dists := make(map[string]int)
	var q []string
	for k := range edges {
		dists[k] = math.MaxInt32
		q = append(q, k)
	}

	dists["AA"] = 0

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return dists[q[i]] < dists[q[j]]
		})

		current := q[0]
		q = q[1:]

		for k := range edges[current] {
			dist := edges[current][k] + dists[current]
			if dist < dists[k] {
				dists[k] = dist
			}
		}
	}

	return dists
}

var (
	up    = vec2{Y: -1}
	down  = vec2{Y: 1}
	left  = vec2{X: -1}
	right = vec2{X: 1}
)

var directions = []vec2{up, down, left, right}

func findEdgesRec(grid map[vec2]rune, portals map[string]*portal, from string) map[string](map[string]int) {
	ans := make(map[string](map[string]int))

	// edges := bfs(grid, portals[from].from)

	return ans
}

func findEdges(grid map[vec2]rune, portals map[string]*portal) map[string](map[string]int) {
	ans := make(map[string](map[string]int))

	for k, from := range portals {
		edges := bfs(grid, from.from)
		edgesTo := bfs(grid, from.to)
		for kt, vt := range edgesTo {
			if vf, ok := edges[kt]; ok {
				if vt < vf {
					edges[kt] = vt
				}
			} else {
				edges[kt] = vt
			}
		}
		ans[k] = edges
	}

	return ans
}

type queue struct {
	pos  vec2
	dist int
	path []vec2
}

func bfs(grid map[vec2]rune, from vec2) map[string]int {
	q := []queue{queue{from, 1, []vec2{from}}}
	visited := make(map[vec2]bool)
	ans := make(map[string]int)

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		visited[current.pos] = true
		for _, dir := range directions {
			next := current.pos.Add(dir)
			if !visited[next] {
				if grid[next] == '.' {
					q = append(q, queue{next, current.dist + 1, append(current.path, next)})
				}
				if isLetter(grid[next]) && current.dist != 1 {
					p, _ := getPortal(grid, next)
					ans[p.name] = current.dist
				}
			}
		}
	}

	return ans
}

func findPortals(grid map[vec2]rune) map[string]*portal {
	ans := make(map[string]*portal)
	dim := utils.GridSize(grid)
	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			pos := vec2{X: x, Y: y}
			val := grid[pos]
			if isLetter(val) {
				if p, ok := getPortal(grid, pos); ok {
					if _, ok := ans[p.name]; ok {
						ans[p.name].to = p.from
					} else {
						ans[p.name] = &portal{p.name, p.from, vec2{}}
					}
				}
			}
		}
	}
	return ans
}

func isLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func getPortal(grid map[vec2]rune, pos vec2) (portal, bool) {
	var builder strings.Builder
	var entry vec2
	for y := pos.Y - 1; y <= pos.Y+1; y++ {
		for x := pos.X - 1; x <= pos.X+1; x++ {
			p := vec2{X: x, Y: y}
			c := grid[p]
			if isLetter(c) {
				builder.WriteRune(c)
			}
			if c == '.' {
				entry = p
			}
		}
	}
	ans := builder.String()
	if len(ans) == 2 && !entry.Eq(vec2{}) {
		return portal{name: ans, from: entry}, true
	}
	return portal{}, false
}
