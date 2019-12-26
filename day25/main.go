package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

var blacklist = map[string]bool{
	"giant electromagnet": true,
	"infinite loop":       true,
	"molten lava":         true,
	"escape pod":          true,
	"photons":             true,
}

var dirs = map[string]vec2{
	"north": vec2{Y: -1},
	"east":  vec2{X: 1},
	"south": vec2{Y: 1},
	"west":  vec2{X: -1},
}

type vec2 = utils.Vector2D

func main() {
	code := utils.ReadIntCode("input.txt")

	input, output := make(chan int, 65536), make(chan int)
	go utils.IntCodeMachine(code, input, output)

	halt := make(chan bool)

	stroutput := make(chan string)
	command := make(chan string, 16)

	go func() {
		var builder strings.Builder
		for x := range output {
			if x < 255 {
				if x == '\n' {
					stroutput <- builder.String()
					builder.Reset()
				} else {
					builder.WriteRune(rune(x))
				}
			} else {
				fmt.Println(x)
			}
		}
		halt <- true
	}()

	go func() {
		for com := range command {
			for _, c := range com {
				input <- int(c)
			}
			input <- '\n'
		}
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			command <- scanner.Text()
		}
	}()

	var currentRoom string
	rooms := make(map[string]bool)
	var mode int
	var (
		items      []string
		inventory  []string
		directions []string
	)

	var combinations [][]string
	var combosIdx int

	var directionToPSF string
	var lastDirection string
	searching := true

	skipCommand := 0

	for out := range stroutput {
		// fmt.Println(out)
		if strings.HasPrefix(out, "==") {
			lastRoom := currentRoom
			currentRoom = strings.Trim(out, "= ")
			if strings.HasPrefix(lastRoom, "Pressure-Sensitive Floor") && strings.HasPrefix(currentRoom, "Security Checkpoint") {
				directionToPSF = lastDirection
			}
			if _, ok := rooms[currentRoom]; !ok {
				rooms[currentRoom] = true
			}
			items = nil
			directions = nil
		}
		if strings.HasPrefix(out, "Doors") {
			mode = 1
		}
		if strings.HasPrefix(out, "Items") {
			mode = 2
		}
		if strings.HasPrefix(out, "- ") {
			option := strings.Trim(out, "- ")
			if mode == 1 {
				directions = append(directions, option)
			}
			if mode == 2 {
				if _, ok := blacklist[option]; !ok {
					items = append(items, option)
				}
			}
		}
		if strings.HasPrefix(out, "Command") {
			if skipCommand > 0 {
				skipCommand--
				continue
			}
			if len(inventory) < 8 && searching {
				if len(items) > 0 {
					command <- fmt.Sprintf("take %s", items[0])
					inventory = append(inventory, items[0])
					items = items[1:]

					if len(inventory) == 8 {
						searching = false
					}
				} else if len(directions) > 0 {
					lastDirection = directions[rand.Intn(len(directions))]
					command <- lastDirection
				}
				if len(inventory) == 8 {
					combinations = generateCombinations(inventory)
				}
			} else if strings.HasPrefix(currentRoom, "Security Checkpoint") && !searching && rooms["Pressure-Sensitive Floor"] {
				// fmt.Println(inventory)
				// reset inventory

				for _, item := range inventory {
					command <- fmt.Sprintf("drop %s", item)
					skipCommand++
				}

				inventory = combinations[combosIdx]
				combosIdx++

				for _, item := range inventory {
					command <- fmt.Sprintf("take %s", item)
					skipCommand++
				}

				// try combo
				command <- directionToPSF
				// skipCommand++
			} else {
				lastDirection = directions[rand.Intn(len(directions))]
				command <- lastDirection
			}
		}
		if strings.HasPrefix(out, "\"Oh, hello") {
			re, _ := regexp.Compile(`\d+`)
			fmt.Println(re.FindAllString(out, -1)[0])
			os.Exit(0)
		}
	}
	<-halt

}

func generateCombinations(items []string) [][]string {
	l := len(items)
	var ans [][]string

	for sub := 1; sub < (1 << l); sub++ {
		var subset []string

		for o := 0; o < l; o++ {
			if (sub>>o)&1 == 1 {
				subset = append(subset, items[o])
			}
		}

		ans = append(ans, subset)
	}
	return ans
}
