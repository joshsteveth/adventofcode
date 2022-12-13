package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

type cmd int

const (
	rock cmd = iota + 1
	paper
	scissor

	lose cmd = 0
	draw cmd = 3
	win  cmd = 6
)

var (
	opponentMap = map[string]cmd{
		"A": rock,
		"B": paper,
		"C": scissor,
	}

	responseMap = map[string]cmd{
		"X": rock,
		"Y": paper,
		"Z": scissor,
	}

	resultMap = map[string]cmd{
		"X": lose,
		"Y": draw,
		"Z": win,
	}

	cmdByResult = map[cmd]map[cmd]cmd{
		rock: map[cmd]cmd{
			lose: scissor,
			win:  paper,
			draw: rock,
		},
		paper: map[cmd]cmd{
			lose: rock,
			win:  scissor,
			draw: paper,
		},
		scissor: map[cmd]cmd{
			lose: paper,
			win:  rock,
			draw: scissor,
		},
	}
)

func parseInput(str string) (opponent, response cmd) {
	s := strings.Split(str, " ")
	return opponentMap[s[0]], responseMap[s[1]]
}

func parseInputPart2(str string) (opponent, response cmd) {
	s := strings.Split(str, " ")
	return opponentMap[s[0]], resultMap[s[1]]
}

func result(opponent, response cmd) (score cmd) {
	switch opponent {
	case rock:
		switch response {
		case rock:
			return draw
		case paper:
			return win
		default:
			return lose
		}
	case paper:
		switch response {
		case rock:
			return lose
		case paper:
			return draw
		default:
			return win
		}
	default:
		switch response {
		case rock:
			return win
		case paper:
			return lose
		default:
			return draw
		}
	}
}

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

	var total int
	for _, inp := range inputs {
		opp, resp := parseInput(inp)
		total += int(result(opp, resp) + resp)
	}

	fmt.Printf("total score: %d\n", total)
}

func star2(inputs []string) {
	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	var total int
	for _, inp := range inputs {
		opp, result := parseInputPart2(inp)
		cmdMap := cmdByResult[opp]
		total += int(result + cmdMap[result])
	}

	fmt.Printf("total score v2: %d\n", total)
}
