package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	input := utils.ReadLines("input.txt")[0]
	fmt.Println(len(input))
	digits := toDigits(input)
	part1 := fft(digits, 100)
	fmt.Println(toInt(part1[:8]))
	digits = repeat(digits, 10000)
	offset := toInt(digits[:7])
	part2 := fft(digits[offset:], 100)
	fmt.Println(toInt(part2[:8]))
}

func repeat(array []int, times int) []int {
	ans := make([]int, len(array)*times)
	for i := 0; i < times; i++ {
		for j, val := range array {
			ans[i*len(array)+j] = val
		}
	}
	return ans
}

func toDigits(input string) []int {
	ans := make([]int, len(input))
	for i, x := range input {
		ans[i] = int(x - '0')
	}
	return ans
}

func toInt(digits []int) int {
	ans := 0
	for _, x := range digits {
		ans *= 10
		ans += x
	}
	return ans
}

func fft(digits []int, phases int) []int {
	pattern := []int{0, 1, 0, -1}
	ans := make([]int, len(digits))
	copy(ans, digits)
	for phase := 0; phase < phases; phase++ {
		fmt.Println(phase)
		new := make([]int, len(digits))
		for i := 0; i < len(digits); i++ {
			for j := 0; j < len(digits); j++ {
				tmp := pattern[((j+1)/(i+1))%4]
				new[i] = new[i] + ans[j]*tmp
			}
			new[i] = utils.Abs(new[i]) % 10
		}
		copy(ans, new)
	}
	return ans
}
