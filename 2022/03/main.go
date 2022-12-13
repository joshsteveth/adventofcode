package main

import (
	"fmt"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

var testInputs = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func main() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)
	//star1(testInputs)
	//star2(testInputs)

	star1(inputs)
	star2(inputs)
}

func runeToPriority(r rune) int {
	i := int(r) - 96
	if i > 0 {
		return i
	}

	// correction for upper case
	return i + 32 + 26
}

func splitStrings(str string) (string, string) {
	l := len(str) / 2
	return str[:l], str[l:]
}

func findCommonBadge(inputs ...string) int {
	mapStuff := func(s string) map[int]any {
		m := make(map[int]any)
		for _, r := range s {
			prio := runeToPriority(r)
			m[prio] = 1
		}
		return m
	}

	findInComon := func(s1, s2 string) map[int]any {
		m2 := mapStuff(s2)
		m := make(map[int]any)
		for _, r := range s1 {
			prio := runeToPriority(r)
			if _, ok := m2[prio]; ok {
				m[prio] = 1
			}
		}
		return m
	}

	commonMap := findInComon(inputs[0], inputs[1])
	for i := 2; i < len(inputs); i++ {
		// eliminate the entries in commonMap by adding more inputs
		m := mapStuff(inputs[i])
		for prio := range commonMap {
			if _, ok := m[prio]; !ok {
				delete(commonMap, prio)
			}
		}
	}
	var prio int
	for v := range commonMap {
		prio = v
	}
	return prio
}

func star1(inputs []string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	var total int
	for _, inp := range inputs {
		s1, s2 := splitStrings(inp)
		// total += (findCommonItem(s1, s2))
		total += findCommonBadge(s1, s2)
	}
	fmt.Printf("[star1] total prio: %d\n", total)
}

func star2(inputs []string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	var (
		total int
		tuple = [3]string{}
	)

	for i, inp := range inputs {
		mod := i % 3
		tuple[mod] = inp

		if mod != 2 {
			continue
		}

		total += findCommonBadge(tuple[0], tuple[1], tuple[2])
	}

	fmt.Printf("[star2] total badge: %d\n", total)
}
