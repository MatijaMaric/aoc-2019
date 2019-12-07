package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	code := utils.IntList(utils.ReadLines("input.txt")[0])
	maxThrust := 0
	for _, permutation := range generatePermutations([]int{5, 6, 7, 8, 9}) {
		thrust := runAmplifiers(code, permutation)
		maxThrust = utils.Max(maxThrust, thrust)
	}
	fmt.Println(maxThrust)
}
