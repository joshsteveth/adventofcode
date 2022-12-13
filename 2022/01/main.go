package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

func main() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)
	star1(inputs)
	star2(inputs)
}

func star1(inputs []string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	var (
		max     int
		tempVal int
	)

	for _, inp := range inputs {
		if inp == "" {
			if tempVal > max {
				max = tempVal
			}
			tempVal = 0
			continue
		}
		i, err := strconv.Atoi(inp)
		util.Must(fmt.Errorf("unable to convert input %s: %w", inp, err))
		tempVal += i
	}

	fmt.Printf("max value: %d\n", max)
}

func star2(inputs []string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	var (
		vals    = []int{0, 0, 0}
		tempVal int
	)

	for _, inp := range inputs {
		if inp == "" {
			lastVal := vals[2]
			if tempVal > lastVal {
				vals[2] = tempVal
				sort.Slice(vals, func(i, j int) bool { return vals[j] < vals[i] })
			}
			tempVal = 0
			continue
		}
		i, err := strconv.Atoi(inp)
		util.Must(fmt.Errorf("unable to convert input %s: %w", inp, err))
		tempVal += i
	}
	fmt.Println(vals)
	fmt.Printf("top 3 combined value: %d\n", vals[0]+vals[1]+vals[2])
}
