package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	mass := utils.ReadNumbers("input.txt")

	var (
		total  int = 0
		total2 int = 0
	)

	for _, mass := range mass {
		total += mass/3 - 2
		for {
			fuel := mass/3 - 2
			if fuel <= 0 {
				break
			}
			total2 += fuel
			mass = fuel
		}
	}

	fmt.Println(total)
	fmt.Println(total2)
}
