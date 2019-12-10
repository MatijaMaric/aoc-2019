package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec2 = utils.Vector2D

func main() {
	asteroidMap := utils.ReadLines("input.txt")
	asteroids := getAsteroids(asteroidMap)

	max := 0
	for asteroid := range asteroids {
		visible := findVisible(asteroids, asteroid)
		max = utils.Max(max, len(visible))
	}

	fmt.Println(max)
}

func getAsteroids(asteroidMap []string) map[vec2]bool {
	ans := make(map[vec2]bool)
	for x, row := range asteroidMap {
		for y, cell := range row {
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
