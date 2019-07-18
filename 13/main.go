package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

const (
	filename = "input.txt"
	interval = time.Millisecond * 500
	verbose  = false
)

const (
	leftArr, rightArr, upArr, downArr = "<", ">", "^", "v"

	horz, vert = "-", "|"

	intersection = "+"

	left, straight, right = "left", "straight", "right"
)

type cart struct {
	form                 string
	nextIntersectionTurn string
	x, y                 int
}

func (c *cart) intersection() string {
	switch c.form {
	case leftArr:

		switch c.nextIntersectionTurn {
		case left:
			c.form = downArr
			c.nextIntersectionTurn = straight
		case straight:
			c.nextIntersectionTurn = right
		case right:
			c.form = upArr
			c.nextIntersectionTurn = left
		}
	case rightArr:

		switch c.nextIntersectionTurn {
		case left:
			c.form = upArr
			c.nextIntersectionTurn = straight
		case straight:
			c.nextIntersectionTurn = right
		case right:
			c.form = downArr
			c.nextIntersectionTurn = left
		}
	case upArr:
		switch c.nextIntersectionTurn {
		case left:
			c.form = leftArr
			c.nextIntersectionTurn = straight
		case straight:
			c.nextIntersectionTurn = right
		case right:
			c.form = rightArr
			c.nextIntersectionTurn = left
		}
	case downArr:
		switch c.nextIntersectionTurn {
		case left:
			c.form = rightArr
			c.nextIntersectionTurn = straight
		case straight:
			c.nextIntersectionTurn = right
		case right:
			c.form = leftArr
			c.nextIntersectionTurn = left
		}
	}

	return c.form
}

func newCart(x, y int, form string) *cart {
	return &cart{
		x:                    x,
		y:                    y,
		form:                 form,
		nextIntersectionTurn: left,
	}
}

func main() {

	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	var oldMap [][]string

	var carts []*cart

	for y, l := range lines {

		var c []string

		for x := range l {
			newChar := string(l[x])
			var isNewCart bool
			switch newChar {
			case leftArr, rightArr:
				isNewCart = true
				newChar = horz
			case upArr, downArr:
				isNewCart = true
				newChar = vert
			}

			if isNewCart {
				carts = append(carts, newCart(x, y, string(l[x])))
			}

			c = append(c, newChar)
		}
		oldMap = append(oldMap, c)
	}

	currentMap := renderCoordinates(oldMap, carts)

	var counter int

	//scanner := bufio.NewScanner(os.Stdin)
	for {
		counter++
		fmt.Printf("%d.iteration...\n", counter)

		// sort by the one on top first
		sort.Slice(carts, func(i, j int) bool {
			ci, cj := carts[i], carts[j]
			if ci.y == cj.y {
				return ci.x < cj.x
			}
			return ci.y < cj.y
		})

		for _, c := range carts {
			fmt.Println("")
			// move all carts 1 step
			var err error
			currentMap, err = c.move(oldMap, currentMap)
			if err != nil {
				fmt.Println(err)
				return
			}

		}

		if verbose {
			for _, coordinate := range currentMap {
				fmt.Println(coordinate)
			}
			time.Sleep(interval)
		}

	}

}

func (c *cart) move(oldMap, maps [][]string) (newMap [][]string, err error) {

	newMap = copyMap(maps)

	newCoord := [2]int{}
	var newChar string

	switch c.form {
	case leftArr:
		newCoord = [2]int{c.x - 1, c.y}

		nextTile := maps[c.y][c.x-1]
		// moving left
		// possible next tile is "-", "\", "/", "+"
		switch nextTile {
		case horz:
			newChar = leftArr
		case `\`:
			newChar = upArr
		case "/":
			newChar = downArr
		case intersection:
			newChar = c.intersection()
		default:
			err = fmt.Errorf("[leftArr] accident happen in %v", newCoord)
		}
	case rightArr:
		newCoord = [2]int{c.x + 1, c.y}

		nextTile := maps[c.y][c.x+1]
		// moving left
		// possible next tile is "-", "\", "/", "+"
		switch nextTile {
		case horz:
			newChar = rightArr
		case `\`:
			newChar = downArr
		case "/":
			newChar = upArr
		case intersection:
			newChar = c.intersection()
		default:
			err = fmt.Errorf("[rightArr] accident happen in %v %s", newCoord, nextTile)
		}
	case upArr:
		newCoord = [2]int{c.x, c.y - 1}

		nextTile := maps[c.y-1][c.x]
		// moving left
		// possible next tile is "|", "\", "/", "+"
		switch nextTile {
		case vert:
			newChar = upArr
		case `\`:
			newChar = leftArr
		case "/":
			newChar = rightArr
		case intersection:
			newChar = c.intersection()
		default:
			err = fmt.Errorf("[upArr] accident happen in %v", newCoord)
		}
	case downArr:
		newCoord = [2]int{c.x, c.y + 1}

		nextTile := maps[c.y+1][c.x]
		// moving left
		// possible next tile is "|", "\", "/", "+"
		switch nextTile {
		case vert:
			newChar = downArr
		case `\`:
			newChar = rightArr
		case "/":
			newChar = leftArr
		case intersection:
			newChar = c.intersection()
		default:
			err = fmt.Errorf("[downArr] accident happen in %v", newCoord)
		}
	}

	if err != nil {
		newMap[newCoord[1]][newCoord[0]] = oldMap[newCoord[1]][newCoord[0]]
		c.x = newCoord[0]
		c.y = newCoord[1]
		return newMap, err

	}

	newMap[c.y][c.x] = oldMap[c.y][c.x]
	newMap[newCoord[1]][newCoord[0]] = newChar

	c.x = newCoord[0]
	c.y = newCoord[1]
	c.form = newChar

	return newMap, nil
}

func renderCoordinates(maps [][]string, carts []*cart) (coordinates [][]string) {

	coordinates = copyMap(maps)

	for _, c := range carts {
		coordinates[c.y][c.x] = c.form
	}

	for _, c := range coordinates {
		fmt.Println(c)
	}

	return
}

func copyMap(maps [][]string) (coordinates [][]string) {

	for i := 0; i < len(maps); i++ {
		ln := len(maps[i])
		coordinates = append(coordinates, make([]string, ln))
		for j := 0; j < ln; j++ {
			coordinates[i][j] = maps[i][j]
		}
	}

	return
}
