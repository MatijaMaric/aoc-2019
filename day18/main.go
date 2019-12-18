package main

import (
	"fmt"
	"math"
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

var chars = "abcdefghijklmnopqrstuvwxyz@"

func main() {
	grid := utils.ReadGrid("input.txt")

	utils.PrintGrid(grid)

	edges := getEdges(grid)
	min := solve(edges)

	fmt.Println(min)
}

func printEdge(e map[rune]edge) {
	for k, v := range e {
		fmt.Printf("%c ", k)
		fmt.Println(v)
	}
}

type node struct {
	val  rune
	keys string
	dist int
}

func bitwiseKeys(keys string) int {
	ans := 0
	for i, c := range chars {
		if strings.ContainsRune(keys, c) {
			ans |= 1 << i
		}
	}
	return ans
}

func solve(edges map[rune](map[rune]edge)) int {
	visited := make(map[rune](map[int]bool))
	for _, c := range chars {
		visited[c] = make(map[int]bool)
	}
	queue := []node{node{'@', "", 0}}

	min := math.MaxInt32

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.val][bitwiseKeys(current.keys)] {
			continue
		}
		visited[current.val][bitwiseKeys(current.keys)] = true

		if len(current.keys) == 26 {
			min = utils.Min(min, current.dist)
			if min < current.dist {
				fmt.Println(min)
			}
		}

		for k, v := range edges[current.val] {
			if !containsAll(v.gates, current.keys) || strings.ContainsRune(current.keys, k) {
				continue
			}
			newKeys := utils.AppendRune(current.keys, k)
			newNode := node{k, newKeys, current.dist + v.dist}
			queue = append(queue, newNode)
		}
	}

	return min
}

func containsAll(gates, keys string) bool {
	for _, c := range gates {
		if !strings.ContainsRune(keys, c) {
			return false
		}
	}
	return true
}

func getEdges(grid map[vec2]rune) map[rune](map[rune]edge) {
	ans := make(map[rune](map[rune]edge))
	for _, node := range chars {
		ans[node] = bfs(grid, getNode(grid, node))
	}
	return ans
}

type edge struct {
	pos   vec2
	gates string
	dist  int
}

func bfs(grid map[vec2]rune, start vec2) map[rune]edge {
	visited := make(map[vec2]bool)
	queue := []edge{edge{start, "", 0}}
	keys := make(map[rune]edge)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visited[current.pos] = true
		val := grid[current.pos]
		if isKey(val) {
			keys[val] = current
		}
		for _, dir := range directions {
			next := edge{current.pos.Add(dir), current.gates, current.dist + 1}
			if !visited[next.pos] && grid[next.pos] != '#' {
				val = grid[next.pos]
				if isGate(val) {
					next.gates = utils.AppendRune(next.gates, val-'A'+'a')
				}
				queue = append(queue, next)
			}
		}
	}
	return keys
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
