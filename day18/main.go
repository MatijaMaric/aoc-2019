package main

import (
	"fmt"
	"math"
	"os"
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

var directions = []vec2{up, right, down, left}

func main() {
	grid := utils.ReadGrid("input.txt")

	utils.PrintGrid(grid)

	start := getNode(grid, '@')

	part1 := findKeys(grid, start, "", 0)

	file, _ := os.Create("out.txt")
	defer file.Close()
	fmt.Fprintln(file, part1)

}

func findKeys(grid map[vec2]rune, start vec2, keys string, dist int) int {
	if len(keys) == 26 {
		return dist
	}
	reachable := bfs(grid, start, keys)
	min := math.MaxInt32
	c := make(chan int, 26)
	for k, v := range reachable {
		newKeys := utils.AppendRune(keys, k)
		go func(p vec2, in string, d int) {
			c <- findKeys(grid, p, in, d)
		}(v.pos, newKeys, dist+v.dist)
	}

	for i := 0; i < len(reachable); i++ {
		min = utils.Min(min, <-c)
	}
	return min
}

type keyPos struct {
	pos vec2

	dist int
}

func bfs(grid map[vec2]rune, start vec2, keys string) map[rune]keyPos {
	visited := make(map[vec2]bool)
	queue := []keyPos{keyPos{start, 0}}
	gates := strings.ToUpper(keys)
	reachable := make(map[rune]keyPos)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visited[current.pos] = true
		val := grid[current.pos]
		if isKey(val) && !strings.ContainsRune(keys, val) {
			reachable[val] = current
		}
		for _, dir := range directions {
			next := keyPos{current.pos.Add(dir), current.dist + 1}
			if !visited[next.pos] && grid[next.pos] != '#' {
				val = grid[next.pos]
				if isGate(val) && !strings.ContainsRune(gates, val) {
					continue
				}
				queue = append(queue, next)
			}
		}
	}
	return reachable
}

func getNode(grid map[vec2]rune, node rune) vec2 {
	for k, v := range grid {
		if v == node {
			return k
		}
	}
	panic("noooo")
}

func isKey(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func isGate(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func sortString(str string) string {
	runes := utils.StrToRunes(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return utils.RunesToStr(runes)
}
