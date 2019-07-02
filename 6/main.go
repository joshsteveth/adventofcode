package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

const filename = "input.txt"

type coordinate struct {
	x, y float64
}

func (c *coordinate) manhattanDistance(c2 *coordinate) float64 {
	return math.Abs(c2.x-c.x) + math.Abs(c2.y-c.y)
}

func (c *coordinate) calculateMinDistance(coordinates []*coordinate) (minDist float64, idx int) {

	minDist = -1
	var indices []int

	for i, c2 := range coordinates {
		d := c.manhattanDistance(c2)

		// 0 can't be beaten, aka the same location
		if d == 0 {
			return 0, i
		}

		if minDist == -1 || d < minDist {
			minDist = d
			indices = []int{i}
			continue
		}

		// append index to indices if the value is the same
		if d == minDist {
			indices = append(indices, i)
			continue
		}

	}

	// there are multiple indices with same minimum distance
	if len(indices) > 1 {
		return minDist, -1
	}

	return minDist, indices[0]
}

var (
	max_X, max_Y float64

	coordinates      []*coordinate
	coordinateValues []float64
)

func newCoordinate(s string) (*coordinate, error) {
	str := strings.Split(strings.Replace(s, " ", "", -1), ",")
	x, err := strconv.ParseFloat(str[0], 64)
	if err != nil {
		return nil, err
	}
	y, err := strconv.ParseFloat(str[1], 64)
	if err != nil {
		return nil, err
	}
	return &coordinate{x, y}, nil
}

func main() {

	printMap := false

	startTime := time.Now()

	defer func(t time.Time) {
		fmt.Println("runtime: ", time.Since(t))
	}(startTime)

	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	coordinates = make([]*coordinate, 0)
	for _, l := range lines {
		c, err := newCoordinate(l)
		if err != nil {
			panic(err)
		}

		if c.x > max_X {
			max_X = c.x
		}

		if c.y > max_Y {
			max_Y = c.y
		}

		coordinates = append(coordinates, c)
	}

	coordinateValues = make([]float64, len(coordinates))

	fmt.Printf("number of coordinates: %d\nmax_x: %v\nmax_y: %v\n",
		len(coordinates), max_X, max_Y)

	outOfCompetition := map[int]struct{}{}

	alphabets := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for y := float64(0); y <= max_Y+1; y++ {
		for x := float64(0); x <= max_X+1; x++ {

			c := &coordinate{x, y}

			_, idx := c.calculateMinDistance(coordinates)

			// we have more than 1 matching indices
			if idx < 0 {
				if printMap {
					fmt.Print(".")
				}
				continue
			}

			if printMap {
				fmt.Print(string(alphabets[idx]))
			}

			coordinateValues[idx]++

			if x == 0 || x == max_X || y == 0 || y == max_Y {
				outOfCompetition[idx] = struct{}{}
			}

		}
		if printMap {
			fmt.Print("\n")
		}
	}

	var max float64

	for i := 0; i < len(coordinates); i++ {
		val := coordinateValues[i]

		var mark string

		if _, ok := outOfCompetition[i]; ok {
			mark = "[x]"
		} else {
			if val > max {
				max = val
			}
		}

		fmt.Printf("%+v : %v %s\n", coordinates[i], val, mark)

	}

	fmt.Printf("Max Val: %v\n", max)

}
