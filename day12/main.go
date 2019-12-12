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

	initialX, initialY, initialZ := splitAxis(moons)

	for i := 0; i < 1000; i++ {
		simulate(moons)
	}
	energy := 0
	for _, moon := range moons {
		energy += moon.position.L1Norm() * moon.velocity.L1Norm()
	}

	fmt.Println(energy)

	moonX := make([]moon, len(initialX))
	moonY := make([]moon, len(initialY))
	moonZ := make([]moon, len(initialZ))
	copy(moonX, initialX)
	copy(moonY, initialY)
	copy(moonZ, initialZ)

	var stepsX, stepsY, stepsZ int
	for {
		simulate(moonX)
		stepsX++
		if compareState(moonX, initialX) {
			break
		}
	}
	for {
		simulate(moonY)
		stepsY++
		if compareState(moonY, initialY) {
			break
		}
	}
	for {
		simulate(moonZ)
		stepsZ++
		if compareState(moonZ, initialZ) {
			break
		}
	}
	steps := utils.Lcm(utils.Lcm(stepsX, stepsY), stepsZ)
	fmt.Println(steps)
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

func splitAxis(moons []moon) (moonX []moon, moonY []moon, moonZ []moon) {
	for i := range moons {
		moonX = append(moonX, moon{moons[i].position.JustX(), moons[i].velocity.JustX()})
		moonY = append(moonY, moon{moons[i].position.JustY(), moons[i].velocity.JustY()})
		moonZ = append(moonZ, moon{moons[i].position.JustZ(), moons[i].velocity.JustZ()})
	}
	return
}
