package main

import (
	"fmt"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

func main() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)

	testStar1()
	star1(inputs[0])

	testStar2()
	star2(inputs[0])
}

func testStar1() {
	for s, want := range map[string]int{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	} {
		got := firstMarker(s, 4)
		if got != want {
			panic(fmt.Sprintf("unexpected result for %s. want %d but got %d", s, want, got))
		}
	}
	fmt.Println("test * OK")
}

func testStar2() {
	for s, want := range map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	} {
		got := firstMarker(s, 14)
		if got != want {
			panic(fmt.Sprintf("unexpected result for %s. want %d but got %d", s, want, got))
		}
	}
	fmt.Println("test ** OK")
}

func star1(input string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)
	fmt.Printf("[*] result: %d\n", firstMarker(input, 4))
}

func star2(input string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)
	fmt.Printf("[**] result: %d\n", firstMarker(input, 14))
}

func firstMarker(s string, n int) int {
	buff := make([]rune, 0, n)

	isUnique := func() bool {
		m := map[rune]any{}
		for _, r := range buff {
			if _, ok := m[r]; ok {
				return false
			}
			m[r] = struct{}{}
		}
		return true
	}

	for i, r := range s {
		if i < n {
			// fill the buffer in the early phase
			buff = append(buff, r)
			continue
		}

		if isUnique() {
			return i
		}

		buff = buff[1:]
		buff = append(buff, r)
	}

	return 0
}
