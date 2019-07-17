package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/joshsteveth/adventofcode/util"
)

const filename string = "input.txt"

func main() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	m := make(map[int]map[int]int)

	for _, l := range lines {
		claim, err := decodeInput(l)
		if err != nil {
			log.Fatalf("failed to decode claim: %v", err)
		}

		populateMap(m, claim)
	}

	var coveredArea int
	for _, cols := range m {
		for _, val := range cols {
			if val > 1 {
				coveredArea++
			}
		}
	}

	fmt.Printf("1st Part: %d square inches of fabric are within two or more claims\n", coveredArea)

	// 2nd part is to get the ID which doesn't overlap
	for _, l := range lines {
		claim, err := decodeInput(l)
		if err != nil {
			log.Fatalf("failed to decode claim: %v", err)
		}

		if doesntOverlap(m, claim) {
			fmt.Printf("2nd Part: claim with ID #%s doesn't overlap with the others\n", claim.id)
			return
		}
	}
	fmt.Printf("2nd Part: every claims overlap")
}

type claim struct {
	id          string
	leftPadding int
	topPadding  int
	width       int
	height      int
}

// decode input from string into claim structure
// e.g. #123 @ 3,2: 5x4 means:
// id = 123
// leftPadding = 3, topPadding = 2
// width = 5, height = 4
func decodeInput(s string) (claim, error) {
	var c claim

	// split by "@" first to get the id
	// trim space and trim prefix "#" from the first element
	at := strings.Split(s, "@")
	if len(at) != 2 {
		return c, fmt.Errorf("invalid format: not splittable by @")
	}
	c.id = strings.TrimPrefix(strings.TrimSpace(at[0]), "#")

	// now we want to split the second element by ":"
	// left side is the paddings, right side is the rectangle size
	colon := strings.Split(at[1], ":")
	if len(colon) != 2 {
		return c, fmt.Errorf("invalid format: not splittable by ':' to determine paddings and size")
	}

	// now we need to determine the padding
	// split again and this time by ","
	// we need to have 2 elements and each of them should be convertable into int
	paddings := strings.Split(strings.TrimSpace(colon[0]), ",")
	if len(paddings) != 2 {
		return c, fmt.Errorf("invalid format: paddings not splittable by ,")
	}
	var err error
	if c.leftPadding, err = strconv.Atoi(paddings[0]); err != nil {
		return c, fmt.Errorf("unable to convert left padding into int: %v", err)
	}
	if c.topPadding, err = strconv.Atoi(paddings[1]); err != nil {
		return c, fmt.Errorf("unable to convert top padding into int: %v", err)
	}

	// finally we get the width and height from the last element
	// it is analogoe to paddings but separated by "x"
	size := strings.Split(strings.TrimSpace(colon[1]), "x")
	if len(size) != 2 {
		return c, fmt.Errorf("invalid format: size not splittable by x")
	}
	if c.width, err = strconv.Atoi(size[0]); err != nil {
		return c, fmt.Errorf("unable to convert width into int: %v", err)
	}
	if c.height, err = strconv.Atoi(size[1]); err != nil {
		return c, fmt.Errorf("unable to convert height into int: %v", err)
	}

	return c, nil
}

func (c claim) getCoordinate() (startRow, endRow, startCol, endCol int) {
	startRow, endRow = c.topPadding, c.topPadding+c.height
	startCol, endCol = c.leftPadding, c.leftPadding+c.width
	return
}

// populateMap populates the map m with values from claim c
// m[1][2] -> means the value of 2nd row and 3rd column of the map
func populateMap(m map[int]map[int]int, c claim) {
	startRow, endRow, startCol, endCol := c.getCoordinate()

	for row := startRow; row < endRow; row++ {
		for col := startCol; col < endCol; col++ {
			// check if map for this row is already available or not
			if _, ok := m[row]; !ok {
				m[row] = make(map[int]int)
			}
			m[row][col]++
		}
	}
}

// doesntOverlap returns whether this claim is overlapping with others or not
// overlapping means that value from map m is > 1
func doesntOverlap(m map[int]map[int]int, c claim) bool {
	startRow, endRow, startCol, endCol := c.getCoordinate()

	for row := startRow; row < endRow; row++ {
		for col := startCol; col < endCol; col++ {
			if m[row][col] > 1 {
				return false
			}
		}
	}

	return true
}
