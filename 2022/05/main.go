package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

// [D]
// [N] [C]
// [Z] [M] [P]
//
//	1   2   3
func example() map[int][]string {
	return map[int][]string{
		1: {"Z", "N"},
		2: {"M", "C", "D"},
		3: {"P"},
	}
}

var exampleInputs = []string{
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

// [G]                 [D] [R]
// [W]         [V]     [C] [T] [M]
// [L]         [P] [Z] [Q] [F] [V]
// [J]         [S] [D] [J] [M] [T] [V]
// [B]     [M] [H] [L] [Z] [J] [B] [S]
// [R] [C] [T] [C] [T] [R] [D] [R] [D]
// [T] [W] [Z] [T] [P] [B] [B] [H] [P]
// [D] [S] [R] [D] [G] [F] [S] [L] [Q]
//
//	1   2   3   4   5   6   7   8   9
func real() map[int][]string {
	return map[int][]string{
		1: {"D", "T", "R", "B", "J", "L", "W", "G"},
		2: {"S", "W", "C"},
		3: {"R", "Z", "T", "M"},
		4: {"D", "T", "C", "H", "S", "P", "V"},
		5: {"G", "P", "T", "L", "D", "Z"},
		6: {"F", "B", "R", "Z", "J", "Q", "C", "D"},
		7: {"S", "B", "D", "J", "M", "F", "T", "R"},
		8: {"L", "H", "R", "B", "T", "V", "M"},
		9: {"Q", "P", "D", "S", "V"},
	}
}

func main() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)
	star1(inputs, real())
	star1(exampleInputs, example())

	star2(inputs, real())
	star2(exampleInputs, example())
}

func star1(inputs []string, m map[int][]string) {
	do(inputs, m, false)
}

func star2(inputs []string, m map[int][]string) {
	do(inputs, m, true)
}

func do(inputs []string, m map[int][]string, keepOrder bool) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	for _, inp := range inputs {
		cmd := newCMD(inp)
		move(m, cmd, keepOrder)
	}

	star := "star 1"
	if keepOrder {
		star = "star 2"
	}

	printResult(m, star)
}

type cmd struct {
	n, from, to int
}

func newCMD(s string) cmd {
	strs := strings.Split(s, " ")
	res := make([]int, 0, 3)
	for _, str := range strs {
		switch str {
		case "move", "from", "to":
			continue
		}
		i, err := strconv.Atoi(str)
		util.Must(err)
		res = append(res, i)
	}
	return cmd{n: res[0], from: res[1], to: res[2]}
}

func move(m map[int][]string, cmd cmd, keepOrder bool) {
	from := m[cmd.from]
	newFromLen := len(from) - cmd.n
	toMove := from[newFromLen:]

	// move it to the new location
	if !keepOrder {
		for i := len(toMove) - 1; i >= 0; i-- {
			m[cmd.to] = append(m[cmd.to], toMove[i])
		}
	} else {
		m[cmd.to] = append(m[cmd.to], toMove...)
	}

	m[cmd.from] = from[:newFromLen]
}

func printResult(m map[int][]string, star string) {
	var res string
	i := 1
	for {
		v, ok := m[i]
		if !ok {
			break
		}
		res += v[len(v)-1]
		i++
	}
	fmt.Printf("[%s] result: %s\n", star, res)
}
