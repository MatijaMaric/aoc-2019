package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	input := utils.ReadLines("input.txt")[0]
	layerRegex := regexp.MustCompile("\\S{150}")
	layers := layerRegex.FindAllString(input, -1)

	fmt.Println(checkSum(layers))

	image := stackLayers(layers)
	rowsRegex := regexp.MustCompile("\\S{25}")
	rows := rowsRegex.FindAllString(image, -1)
	for _, row := range rows {
		pretty := strings.ReplaceAll(row, "0", " ")
		pretty = strings.ReplaceAll(pretty, "1", "#")
		fmt.Println(pretty)
	}
}

func checkSum(layers []string) int {
	minZeros := math.MaxInt32
	minZerosIdx := 0
	for i, layer := range layers {
		zeros := strings.Count(layer, "0")
		if zeros < minZeros {
			minZeros = zeros
			minZerosIdx = i
		}
	}

	ones := strings.Count(layers[minZerosIdx], "1")
	twos := strings.Count(layers[minZerosIdx], "2")
	return ones * twos
}

func stackLayers(layers []string) string {
	ans := strings.Repeat("2", len(layers[0]))
	for _, layer := range layers {
		var builder strings.Builder
		for i, pixel := range layer {
			if ans[i] == '2' {
				builder.WriteRune(pixel)
			} else {
				builder.WriteByte(ans[i])
			}
		}
		ans = builder.String()
	}
	return ans
}
