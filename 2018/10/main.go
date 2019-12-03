package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/joshsteveth/adventofcode/util"
)

const filename = "input.txt"

type coord struct {
	x, y, vx, vy int
}

func (c *coord) updateCoordinate() {
	c.x = c.x + c.vx
	c.y = c.y + c.vy
}

var (
	inputs []*coord
)

func init() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	for _, l := range lines {

		inputs = append(inputs, getCoordinatesAndVelocity(l))
	}
}

func main() {

	fmt.Println("initial input")
	draw(inputs)

	var counter int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counter++
		fmt.Printf("after %d second(s)\n", counter)

		for _, c := range inputs {
			c.updateCoordinate()
		}

		if timeSkip, err := draw(inputs); err != nil {
			fmt.Println(err)

			for i := 0; i < timeSkip; i++ {
				counter++
				for _, c := range inputs {
					c.updateCoordinate()
				}
			}
		}

	}

}

func draw(inputs []*coord) (int, error) {

	const thresholdX, thresholdY = 200, 200

	sort.Slice(inputs, func(i, j int) bool { return inputs[i].x < inputs[j].x })

	minX, maxX := inputs[0].x, inputs[len(inputs)-1].x

	sort.Slice(inputs, func(i, j int) bool { return inputs[i].y < inputs[j].y })

	minY, maxY := inputs[0].y, inputs[len(inputs)-1].y

	rangeX := maxX - minX
	rangeY := maxY - minY

	if rangeX > thresholdX || rangeY > thresholdY {

		timeSkip := 1000
		if rangeX < 6000 {
			timeSkip = 10
		}

		return timeSkip, fmt.Errorf("abort drawing, range too big: rangeX: %d rangeY: %d\n", rangeX, rangeY)
	}

	cart := make([][]int, 0)

	for i := 0; i <= rangeX; i++ {
		cart = append(cart, make([]int, rangeY+1))
	}

	for _, inp := range inputs {
		x := inp.x - minX
		y := inp.y - minY

		cart[x][y] = 1
	}

	fmt.Println("")

	for y := 0; y <= rangeY; y++ {
		for x := 0; x <= rangeX; x++ {
			v := cart[x][y]

			if v == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
	return 0, nil
}

func getCoordinatesAndVelocity(s string) *coord {

	s = strings.ReplaceAll(s, " ", "")
	ss := strings.Split(s, ">velocity=<")
	pos, vel := ss[0], ss[1]

	parseNums := func(s string) (x, y int) {
		xy := strings.Split(s, ",")
		x, _ = strconv.Atoi(xy[0])
		y, _ = strconv.Atoi(xy[1])
		return
	}

	x, y := parseNums(strings.TrimPrefix(pos, "position=<"))
	vx, vy := parseNums(strings.TrimSuffix(vel, ">"))

	return &coord{x, y, vx, vy}
}

/*
after 10577 second(s)

######....##....######..#....#....##.......###...####....####.
.....#...#..#...#.......#...#....#..#.......#...#....#..#....#
.....#..#....#..#.......#..#....#....#......#...#.......#.....
....#...#....#..#.......#.#.....#....#......#...#.......#.....
...#....#....#..#####...##......#....#......#...#.......#.....
..#.....######..#.......##......######......#...#..###..#.....
.#......#....#..#.......#.#.....#....#......#...#....#..#.....
#.......#....#..#.......#..#....#....#..#...#...#....#..#.....
#.......#....#..#.......#...#...#....#..#...#...#...##..#....#
######..#....#..######..#....#..#....#...###.....###.#...####.
*/
