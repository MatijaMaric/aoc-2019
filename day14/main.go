package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2019/utils"
)

type ingredient struct {
	name   string
	amount int
}

type recipe struct {
	result      string
	amount      int
	ingredients []ingredient
}

func main() {
	input := utils.ReadLines("input.txt")
	recipes := getRecipes(input)
	fmt.Println(part1(recipes, 1))
	fmt.Println(part2(recipes))
}

func part1(recipes map[string]recipe, amount int) int {
	stock := make(map[string]int)
	ore := makeIngredient(recipes, stock, ingredient{"FUEL", amount})
	return ore
}

func part2(recipes map[string]recipe) int {
	low := 0
	high := 1000000000000
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if part1(recipes, mid) < 1000000000000 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return mid
}

func makeIngredient(recipes map[string]recipe, stock map[string]int, what ingredient) int {
	if what.name == "ORE" {
		return what.amount
	}
	fromStock := utils.Min(stock[what.name], what.amount)
	what.amount -= fromStock
	stock[what.name] -= fromStock

	factor := (what.amount + recipes[what.name].amount - 1) / recipes[what.name].amount
	stock[what.name] += factor*recipes[what.name].amount - what.amount

	var ans int

	for _, what := range recipes[what.name].ingredients {
		ans += makeIngredient(recipes, stock, ingredient{what.name, what.amount * factor})
	}
	return ans
}

func getRecipes(input []string) map[string]recipe {
	ans := make(map[string]recipe)
	for _, line := range input {
		split := strings.Split(line, " => ")
		result := parseIngredient(split[1])
		split = strings.Split(split[0], ", ")
		var ingredients []ingredient
		for _, ingredient := range split {
			ingredients = append(ingredients, parseIngredient(ingredient))
		}
		ans[result.name] = recipe{result: result.name, amount: result.amount, ingredients: ingredients}
	}
	return ans
}

func parseIngredient(text string) ingredient {
	split := strings.Split(text, " ")
	return ingredient{split[1], utils.ToInt(split[0])}
}
