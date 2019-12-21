package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	code := utils.ReadIntCode("input.txt")

	var vm utils.AsciiCode
	vm.Init(code)

	instructions := strings.Join(utils.ReadLines("part2.spring"), "\n")
	vm.WriteLn(instructions)

	output := vm.Flush()
	fmt.Println(output)

}
