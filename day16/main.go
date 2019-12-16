package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	input := utils.ReadLines("input.txt")[0]
	digits := toDigits(input)
	{
		part1 := make([]int, len(digits))
		copy(part1, digits)
		for i := 0; i < 100; i++ {
			part1 = fft(part1)
		}
		fmt.Println(toInt(part1[:8]))

	}
	{
		digits = repeat(digits, 10000)
		offset := toInt(digits[:7])
		digits = digits[offset:]
		part2 := make([]int, len(digits))
		copy(part2, digits)
		// fmt.Println(len(part2))
		// for i := 0; i < 100; i++ {
		// 	fmt.Println(i)
		// 	part2 = fft(part2)
		// }
		// all ones probably fuck it
		for i := 0; i < 100; i++ {
			for j := len(part2) - 2; j >= 0; j-- {
				part2[j] += part2[j+1]
				part2[j] %= 10
			}
		}
		fmt.Println(toInt(part2[:8]))
	}
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

func fft(digits []int) []int {
	pattern := []int{0, 1, 0, -1}
	ans := make([]int, len(digits))
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			tmp := pattern[((j+1)/(i+1))%4]
			ans[i] = ans[i] + digits[j]*tmp
		}
		ans[i] = utils.Abs(ans[i]) % 10
	}
	return ans
}
