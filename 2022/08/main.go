package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

var example = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

const (
	exampleResult    = 21
	exampleMaxScenic = 8
)

type (
	row  = []int
	grid = []row
)

func main() {
	example1()
	star1()
	example2()
	star2()
}

func example1() {
	grid := newGrid(example)
	got := visible(grid)

	if got != exampleResult {
		panic(fmt.Sprintf("unexpected result. want %d but got %d", exampleResult, got))
	}

	fmt.Println("example 1 successfully executed")
}

func star1() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	grid := newGrid(inputs)
	got := visible(grid)

	fmt.Printf("[*] result: %d\n", got)
}

func example2() {
	grid := newGrid(example)
	got := maxSceneryVal(grid)

	if got != exampleMaxScenic {
		panic(fmt.Sprintf("unexpected result. want %d but got %d", exampleMaxScenic, got))
	}

	fmt.Println("example 2 successfully executed")
}

func star2() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	grid := newGrid(inputs)
	got := maxSceneryVal(grid)

	fmt.Printf("[**] result: %d\n", got)
}

func newGrid(inputs []string) grid {
	res := make([]row, 0, len(inputs))

	for _, input := range inputs {
		row := make([]int, 0, len(input))
		for _, r := range input {
			n, err := strconv.Atoi(string(r))
			util.Must(err)
			row = append(row, n)
		}
		res = append(res, row)
	}
	return res
}

func visible(g grid) int {
	var total int

	check := func(min, max, c, val int, isRow bool) bool {
		for i := min; i < max; i++ {
			var comparison int
			if isRow {
				comparison = g[c][i]
			} else {
				comparison = g[i][c]
			}

			if comparison >= val {
				return false
			}
		}
		return true
	}

	for i, row := range g {
		for j, val := range row {
			// the edges
			if i == 0 || j == 0 || i == len(g)-1 || j == len(row)-1 {
				total++
				continue
			}

			// check left
			if check(0, j, i, val, true) {
				total++
				continue
			}

			// check right
			if check(j+1, len(row), i, val, true) {
				total++
				continue
			}

			// check top
			if check(0, i, j, val, false) {
				total++
				continue
			}

			// check bottom
			if check(i+1, len(g), j, val, false) {
				total++
				continue
			}
		}
	}
	return total
}

func maxSceneryVal(g grid) int {
	var maxVal int

	nTreesVisible := func(min, max, c, val int, isRow, reversed bool) int {
		var (
			total int
		)
		if reversed {
			for i := max - 1; i >= min; i-- {
				total++
				var comparison int
				if isRow {
					comparison = g[c][i]
				} else {
					comparison = g[i][c]
				}
				if val <= comparison {
					break
				}
			}
		} else {
			for i := min; i < max; i++ {
				total++
				var comparison int
				if isRow {
					comparison = g[c][i]
				} else {
					comparison = g[i][c]
				}
				if val <= comparison {
					break
				}
			}
		}
		return total
	}

	for i, row := range g {
		for j, val := range row {
			// the edges are obsolete here
			if i == 0 || j == 0 || i == len(g)-1 || j == len(row)-1 {
				continue
			}

			left := nTreesVisible(0, j, i, val, true, true)
			right := nTreesVisible(j+1, len(row), i, val, true, false)
			top := nTreesVisible(0, i, j, val, false, true)
			bot := nTreesVisible(i+1, len(g), j, val, false, false)

			total := left * right * top * bot

			if total > maxVal {
				maxVal = total
			}
		}
	}
	return maxVal
}
