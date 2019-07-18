package main

import (
	"fmt"
	"strings"

	"github.com/joshsteveth/adventofcode/util"
)

const (
	//input         = "#..#.#..##......###...###"
	//filename      = "example.txt"
	input         = "##..##....#.#.####........##.#.#####.##..#.#..#.#...##.#####.###.##...#....##....#..###.#...#.#.#.#"
	filename      = "input.txt"
	numGeneration = 50000000000
)

type code struct {
	pattern   string
	willBloom bool
}

var codes []*code

func main() {

	const thresholdGeneration = 150

	if numGeneration > thresholdGeneration {
		// we just need to add 80 per new generation
		diff := numGeneration - thresholdGeneration
		fmt.Printf("Final score: %d\n", 12000+diff*80)
		return
	}

	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	paddings := 0

	for _, l := range lines {
		code := decodeLine(l)

		// count number of ..... in a row to determinde the padding size
		var p int
		for _, c := range code.pattern {
			if string(c) == "#" {
				break
			}
			p++
		}

		if p > paddings {
			paddings = p
		}

		codes = append(codes, code)
	}

	plants := input

	// for the sake of convenience, let's add more dots as much as the number of numGenerations
	for i := 0; i < numGeneration; i++ {
		plants += ".."
	}

	var offset, score int

	for x := 0; x < numGeneration; x++ {

		var offs int
		plants, offs = addPadding(plants, paddings)
		offset += offs

		nextUpdate := make([]string, len(plants))

		for i := 2; i < len(plants)-2; i++ {
			nextUpdate[i] = checkAllPatterns(plants[i-2:i+3], codes)
		}

		plants = ""
		for _, u := range nextUpdate {
			if u == "#" {
				plants += u
				continue
			}
			plants += "."
		}
		score = 0
		for i := range plants {
			if plants[i] == '.' {
				continue
			}
			score += (i - offset)
		}
		fmt.Println(plants, score)
	}

	fmt.Printf("Final score: %d\n", score)

}

func checkAllPatterns(s string, codes []*code) string {
	for _, c := range codes {
		if s == c.pattern {
			if c.willBloom {
				return "#"
			}
			return ""
		}
	}
	return ""
}

func addPadding(plants string, padding int) (string, int) {

	// make sure that each "#" has at least 4 "." before it
	for i := 0; i < len(plants); i++ {

		if plants[i] != '#' {
			// we need 1 less padding
			padding--
			continue
		}

		if padding <= 0 {
			return plants, 0
		}

		var p string
		for j := 0; j < padding; j++ {
			p += "."
		}

		return p + plants, padding

	}

	return plants, 0

}

func willBloom(s string) bool {
	if s == "#" {
		return true
	}
	return false
}

func decodeLine(s string) *code {
	str := strings.Split(s, " => ")

	return &code{
		pattern:   str[0],
		willBloom: willBloom(str[1]),
	}
}
