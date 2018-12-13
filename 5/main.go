package main

import (
	"fmt"
	"strings"

	"github.com/joshsteveth/adventofcode/util"
)

const filename = "input.txt"

func main() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	text := lines[0]

	fmt.Printf("initial length: %d\n", len(text))

	text, _ = removeReactable(text, true)

	fmt.Printf("1. Part:\nafter changed length: %d\n", len(text))

	alphabetMap := map[string]int{}
	fmt.Println("2. Part")
	for i := 0; i < 26; i++ {
		a := toStr(i)

		newTxt := text
		newTxt = strings.Replace(newTxt, a, "", -1)
		newTxt = strings.Replace(newTxt, strings.ToUpper(a), "", -1)

		newTxt, _ = removeReactable(newTxt, true)
		fmt.Printf("new length after removing %s: %d\n", a, len(newTxt))

		alphabetMap[a] = len(newTxt)
	}

	lowest := len(text)
	lowestAlpha := ""
	for alpha, val := range alphabetMap {
		if val < lowest {
			lowest = val
			lowestAlpha = alpha
		}
	}

	fmt.Printf("lowest length is achieved by removing %s: %d\n", lowestAlpha, lowest)
}

// capital is actually "lesser" than lower cases
func isCapital(s string) bool {
	if s < "a" {
		return true
	}
	return false
}

// extend a letter by its counterpart
// e.g. : a -> aA, A -> Aa
func extendWord(s string) string {
	if isCapital(s) {
		return s + strings.ToLower(s)
	}
	return s + strings.ToUpper(s)
}

func removeReactable(text string, b bool) (string, bool) {
	if !b {
		return text, b
	}

	for i := 0; i < len(text)-1; i++ {
		// let's see if extended word actually appears on the text
		s := extendWord(text[i : i+1])

		if s == text[i:i+2] {
			// then they should be removed
			return removeReactable(strings.Replace(text, s, "", 1), true)
		}
	}
	return text, false
}

// 0 -> "a"; 25 -> "z"
func toStr(i int) string {
	return string('a' + i)
}
