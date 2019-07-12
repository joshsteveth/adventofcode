package main

import (
	"fmt"
	"sort"
)

const (
	numPlayer = 403
	maxMarble = 71920
)

func main() {

	marbles := []int{0}
	players := make([]int, numPlayer)

	currentMarble := 0

	var activePlayer int

	for p := 1; p <= maxMarble; p++ {
		idx := getIndex(marbles, currentMarble)
		//fmt.Printf("current marble %d with idx %d\n", currentMarble, idx)

		if p%23 == 0 {

			removedIdx := idx - 7

			if removedIdx < 0 {
				removedIdx = len(marbles) + removedIdx
			}

			removedScore := marbles[removedIdx]
			players[activePlayer] += (p + removedScore)

			marbles = removeMarble(marbles, removedIdx)

			currentMarble = marbles[removedIdx]

		} else {
			marbles = addMarble(marbles, p, idx)
			currentMarble = p
		}

		activePlayer++
		if activePlayer >= numPlayer {
			activePlayer = 0
		}

		//fmt.Println(marbles)

	}

	//fmt.Println(players)

	sort.Ints(players)
	fmt.Printf("Highest score: %d\n", players[len(players)-1])
}

func removeMarble(m []int, idx int) []int {
	return append(m[:idx], m[idx+1:]...)
}

func addMarble(m []int, point, currentMarbleIdx int) []int {

	maxLen := len(m)

	if currentMarbleIdx+1 >= maxLen {
		//fmt.Println("on the edge")
		return append([]int{0, point}, m[1:]...)
	}

	// e.g. m = [0  8  4  2  5  1  6  3  7 ]
	leftSide := m[:currentMarbleIdx+2]  // [0  8  4]
	rightSide := m[currentMarbleIdx+2:] // [2  5  1  6  3  7]
	//fmt.Printf("left side %v\nright side: %v\n", leftSide, rightSide)

	return append(leftSide, append([]int{point}, rightSide...)...)
}

func getIndex(m []int, i int) int {
	for idx, marble := range m {
		if marble == i {
			return idx
		}
	}

	return 0
}
