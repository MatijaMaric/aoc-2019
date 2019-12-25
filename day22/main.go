package main

import (
	"regexp"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type Deck []int

func main() {
	instructions := utils.ReadLines("input.txt")

	deck := generateDeck(10007)
	shuffle(deck, instructions)
}

func generateDeck(n int) []int {
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = i
	}
	return ans
}

func shuffle(deck Deck, instructions []string) {
	re := regexp.MustCompile("^([a-z|\\s]+)([-*\\d]*)$")
	for _, instruction := range instructions {
		matches := re.FindAllStringSubmatch(instruction, -1)[0][1:]
		switch matches[0] {
		case "deal into new stack":
			deck.reverse()
		case "cut ":
			val := utils.ToInt(matches[1])
			deck.cut(val)
		case "deal with increment ":
			val := utils.ToInt(matches[1])
			deck.increment(val)
		}
	}
}

func (deck *Deck) reverse() {
	old := *deck
	l := len(old)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		old[i], old[j] = old[j], old[i]
	}
	*deck = old
}

func (deck *Deck) cut(n int) {
	old := *deck
	old = append(old[n:], old[:n]...)
	*deck = old
}

func (deck *Deck) increment(n int) {
	old := *deck
	l := len(old)
	new := make([]int, l)
	for i, newI := 0, 0; i < l; i, newI = i+1, ((i+1)*n)%l {
		new[newI] = old[i]
	}
	*deck = new
}
