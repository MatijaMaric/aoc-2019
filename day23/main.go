package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2019/utils"
)

func main() {
	code := utils.ReadIntCode("input.txt")
	ans := part1(code)
	fmt.Println(ans)
}

type packet struct {
	x, y int
}

func part1(code []int) int {
	queue := make([]chan packet, 50)
	for i := 0; i < 50; i++ {
		queue[i] = make(chan packet, 1024)
	}

	ans := make(chan int)
	for i := 0; i < 50; i++ {
		input, output := make(chan int), make(chan int)
		go utils.IntCodeMachine(code, input, output)
		input <- i

		go func(i int) {
			for p := range queue[i] {
				input <- p.x
				input <- p.y
			}
		}(i)

		go func() {
			for dest := range output {
				x, y := <-output, <-output
				fmt.Println(x, y)
				if dest >= 0 && dest < 50 {
					queue[dest] <- packet{x, y}
				} else {
					ans <- y
				}
			}
		}()
	}
	return <-ans
}
