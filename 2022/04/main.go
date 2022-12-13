package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

var exampleInputs = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func main() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)
	star1(inputs)
	star1(exampleInputs)
	star2(inputs)
}

func star1(inputs []string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)
	var total int
	for _, inp := range inputs {
		a1, a2 := newAssignments(inp)
		if fullyContains(a1, a2) {
			total++
		}
	}
	fmt.Printf("[star 1] total pairs: %d\n", total)
}

func star2(inputs []string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)
	var total int
	for _, inp := range inputs {
		a1, a2 := newAssignments(inp)
		if overlaps(a1, a2) {
			total++
		}
	}
	fmt.Printf("[star 2] total pairs: %d\n", total)
}

type assignment struct {
	low, high int
}

func newAssignments(s string) (assignment, assignment) {
	strs := strings.Split(s, ",")

	newAssignment := func(str string) assignment {
		ss := strings.Split(str, "-")
		low, err := strconv.Atoi(ss[0])
		util.Must(err)
		high, err := strconv.Atoi(ss[1])
		util.Must(err)
		return assignment{low: low, high: high}
	}

	return newAssignment(strs[0]), newAssignment(strs[1])
}

func fullyContains(a1, a2 assignment) bool {
	as := []assignment{a1, a2}
	// first element should have the lower low
	sort.Slice(as, func(i, j int) bool { return as[i].low <= as[j].low })

	first, second := as[0], as[1]

	if first.low == second.low {
		return true
	}

	return second.high <= first.high
}

func overlaps(a1, a2 assignment) bool {
	as := []assignment{a1, a2}
	// first element should have the lower low
	sort.Slice(as, func(i, j int) bool { return as[i].low <= as[j].low })

	first, second := as[0], as[1]

	return first.high >= second.low
}
