package main

import (
	"fmt"
	"regexp"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type vec3 = utils.Vector3D

type moon struct {
	position vec3
	velocity vec3
}

func main() {
	re := regexp.MustCompile(`<x=(-*\d*), y=(-*\d*), z=(-*\d*)>`)
	lines := utils.ReadLines("input.txt")

	var moons []moon

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)[1:]
		moonPosition := vec3{
			X: utils.ToInt(matches[0]),
			Y: utils.ToInt(matches[1]),
			Z: utils.ToInt(matches[2]),
		}
		moons = append(moons, moon{position: moonPosition})
	}

	initialState := make([]moon, len(moons))
	copy(initialState, moons)

	for i := 0; i < 1000; i++ {
		simulate(moons)
	}
	energy := 0
	for _, moon := range moons {
		energy += moon.position.L1Norm() * moon.velocity.L1Norm()
	}

	fmt.Println(energy)

	moons = initialState
	initialState = make([]moon, len(moons))
	copy(initialState, moons)

	count := 0
	for {
		simulate(moons)
		count++
		if compareState(moons, initialState) {
			break
		}
	}
	fmt.Println(count)
}

func simulate(moons []moon) {
	// apply gravity
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			moons[i].velocity = moons[i].velocity.Add(moons[j].position.Subtract(moons[i].position).Sgn())
			moons[j].velocity = moons[j].velocity.Add(moons[i].position.Subtract(moons[j].position).Sgn())
		}
	}

	// apply velocity
	for i := range moons {
		moons[i].position = moons[i].position.Add(moons[i].velocity)
	}
}

func compareState(a, b []moon) bool {
	for i := 0; i < len(a); i++ {
		if !a[i].position.Eq(b[i].position) {
			return false
		}
		if !a[i].velocity.Eq(b[i].velocity) {
			return false
		}
	}
	return true
}
