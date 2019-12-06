package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	orbits := utils.ReadLines("input.txt")
	orbitMap := make(map[string][]string)
	fullLookup := make(map[string][]string)
	allPlanets := make(map[string]bool)
	planetsInOrbit := make(map[string]bool)

	for _, orbit := range orbits {
		planets := strings.Split(orbit, ")")
		orbitMap[planets[0]] = append(orbitMap[planets[0]], planets[1])
		fullLookup[planets[0]] = append(fullLookup[planets[0]], planets[1])
		fullLookup[planets[1]] = append(fullLookup[planets[1]], planets[0])
		allPlanets[planets[0]] = true
		allPlanets[planets[1]] = true
		planetsInOrbit[planets[1]] = true
	}

	var rootPlanet string

	for k := range allPlanets {
		if !planetsInOrbit[k] {
			rootPlanet = k
			break
		}
	}

	count := countOrbits(orbitMap, rootPlanet, 0)
	fmt.Println(count)

	shortest := bfs(fullLookup, "YOU", "SAN")
	fmt.Println(shortest)
}

func countOrbits(orbits map[string][]string, root string, depth int) int {
	planets := orbits[root]
	ans := 0
	for _, planet := range planets {
		ans += countOrbits(orbits, planet, depth+1)
	}

	return depth + ans
}

type node struct {
	planet string
	dist   int
}

func bfs(orbits map[string][]string, from string, to string) int {
	visited := make(map[string]bool)
	var queue = []node{
		node{
			planet: from,
			dist:   0,
		},
	}

	for {
		current := queue[0]
		queue = queue[1:]
		visited[current.planet] = true
		if current.planet == to {
			return current.dist - 2
		}
		for _, planet := range orbits[current.planet] {
			if !visited[planet] {
				queue = append(queue, node{planet: planet, dist: current.dist + 1})
			}
		}
	}
}
