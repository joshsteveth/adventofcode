package main

import (
	"fmt"
)

const (
	numRecipe  = 10
	scoreAfter = 440231

	verbose = false
)

var (
	score = []int{4, 4, 0, 2, 3, 1}
)

func main() {

	recipes := []int{3, 7}

	elfOne, elfTwo := 0, 1

	// scanner := bufio.NewScanner(os.Stdin)

	for {

		if verbose {
			printRecipes(recipes, elfOne, elfTwo)
		}

		// sum value of both elves
		valOne, valTwo := recipes[elfOne], recipes[elfTwo]
		recipes = addSum(recipes, valOne+valTwo)

		// first task
		/*
			if len(recipes) >= numRecipe+scoreAfter {
				result := recipes[scoreAfter : scoreAfter+numRecipe]
				fmt.Printf("Result: %v\n", result)
				return
			}
		*/

		// update idx of both elves
		elfOne = (elfOne + valOne + 1) % len(recipes)
		elfTwo = (elfTwo + valTwo + 1) % len(recipes)

		// second task
		offset := 0
		numMatch := 0

		for i := 0; i < len(score)+1; i++ { // + 1 since it's possible to add 1 or 2 new recipes

			// check if last one is the same or not
			curScore := score[len(score)-1-i+offset]
			curRecipe := recipes[len(recipes)-1-i]

			if verbose {
				fmt.Printf("%v cur score: %d cur recipe: %d\n", recipes, curScore, curRecipe)
			}

			if curScore != curRecipe {

				// bummer..
				// but let's check next one if it's still the first one
				if i == 0 {
					offset = 1
					continue
				}

				goto next
			}
			numMatch++
			if numMatch == len(score) {
				break
			}
		}

		// we have the winner!
		fmt.Printf("Result: %d\n", len(recipes)-len(score)-offset)
		return

	next:
	}

}

func addSum(recipes []int, sum int) (newRecipes []int) {

	if sum/10 == 0 {
		// 1 digit sum
		return append(recipes, sum)
	}

	return append(recipes, []int{1, sum % 10}...)
}

func printRecipes(recipes []int, elfOne, elfTwo int) {
	const border = "===================="
	fmt.Println(border)
	defer fmt.Printf("\n%s\n", border)

	for i, recipe := range recipes {
		var format string
		switch i {
		case elfOne:
			format = "(%d) "
		case elfTwo:
			format = "[%d] "
		default:
			format = "%d "
		}
		fmt.Printf(format, recipe)
	}
}
