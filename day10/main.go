package main

import (
	"fmt"
	"sort"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

func main() {
	asteroidMap := utils.ReadLines("input.txt")
	asteroids := getAsteroids(asteroidMap)

	max := 0
	var base vec2
	for asteroid := range asteroids {
		visible := findVisible(asteroids, asteroid)
		if len(visible) > max {
			max = len(visible)
			base = asteroid
		}
	}

	vaporize := vaporizeSort(asteroids, base)

	fmt.Println(max)
	fmt.Println(vaporize[199].X*100 + vaporize[199].Y)
}

func getAsteroids(asteroidMap []string) map[vec2]bool {
	ans := make(map[vec2]bool)
	for y, row := range asteroidMap {
		for x, cell := range row {
			if cell == '#' {
				ans[vec2{X: x, Y: y}] = true
			}
		}
	}
	return ans
}

func findVisible(asteroids map[vec2]bool, base vec2) []vec2 {
	var ans []vec2

	for asteroid := range asteroids {
		if asteroid.Eq(base) {
			continue
		}
		direction := asteroid.Subtract(base).IntNorm()
		collides := false
		for current := base.Add(direction); !current.Eq(asteroid); current = current.Add(direction) {
			if asteroids[current] {
				collides = true
				break
			}
		}

		if !collides {
			ans = append(ans, asteroid)
		}
	}
	return ans
}

func vaporizeSort(asteroids map[vec2]bool, base vec2) []vec2 {
	var ans []vec2

	for len(asteroids) > 1 {
		visible := findVisible(asteroids, base)
		sort.Slice(visible, func(i, j int) bool {
			angleA := visible[i].Subtract(base).Angle()
			angleB := visible[j].Subtract(base).Angle()
			return angleA < angleB
		})
		fmt.Println(visible)
		ans = append(ans, visible...)
		for _, asteroid := range visible {
			delete(asteroids, asteroid)
		}
	}
	return ans
}
