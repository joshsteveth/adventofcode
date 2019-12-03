// --- Day 3: Crossed Wires ---
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

// The gravity assist was successful, and you're well on your way to the Venus refuelling station.
// During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.

// Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid.
// You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

// The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port.
// Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start,
// this point does not count, nor does a wire count as crossing with itself.

// For example, if the first wire's path is R8,U5,L5,D3, then starting from the central port (o), it goes right 8, up 5, left 5, and finally down 3:

// ...........
// ...........
// ...........
// ....+----+.
// ....|....|.
// ....|....|.
// ....|....|.
// .........|.
// .o-------+.
// ...........

// Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:

// ...........
// .+-----+...
// .|.....|...
// .|..+--X-+.
// .|..|..|.|.
// .|.-X--+.|.
// .|..|....|.
// .|.......|.
// .o-------+.
// ...........

// These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 3 + 3 = 6.

// Here are a few more examples:

//     R75,D30,R83,U83,L12,D49,R71,U7,L72
//     U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
//     R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
//     U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135

// What is the Manhattan distance from the central port to the closest intersection?

type direction string

const (
	right direction = "R"
	left  direction = "L"
	up    direction = "U"
	down  direction = "D"
)

type move struct {
	dir direction
	val int
}

type (
	coord struct{ x, y int }

	// cRange is coordinate range
	cRange struct{ start, end int }

	movement struct{ x, y cRange }
)

func (c coord) manhattanDistance() int {
	if c.x < 0 {
		c.x = -c.x
	}
	if c.y < 0 {
		c.y = -c.y
	}
	return c.x + c.y
}

func (c coord) isIntersected(mov movement) bool {

	if c.x == 0 && c.y == 0 {
		return false
	}

	mov.ascending()

	if c.x < mov.x.start ||
		c.x > mov.x.end ||
		c.y < mov.y.start ||
		c.y > mov.y.end {
		return false
	}

	return true

}

func (c *cRange) ascending() {
	if c.start > c.end {
		c.start, c.end = c.end, c.start
	}
}

func (m *movement) ascending() {
	m.x.ascending()
	m.y.ascending()
}

func (m *movement) steps() int {
	m.ascending()
	return (m.x.end - m.x.start) + (m.y.end - m.x.start)
}

func computeCableMovements(s string) []movement {

	strs := strings.Split(s, ",")
	movs := make([]movement, len(strs))

	cur := &coord{}

	for i, str := range strs {
		m := direction(str[0])
		x, err := strconv.Atoi(str[1:])
		util.Must(err)
		movs[i] = newMovement(cur, move{m, x})
	}

	return movs
}

func newMovement(c *coord, m move) movement {

	// mv is default movement if
	// it doesn't move at all
	mv := movement{
		x: cRange{c.x, c.x},
		y: cRange{c.y, c.y},
	}

	switch m.dir {
	case right:
		mv.x = cRange{c.x, c.x + m.val}
		c.x += m.val
	case left:
		mv.x = cRange{c.x, c.x - m.val}
		c.x -= m.val
	case up:
		mv.y = cRange{c.y, c.y + m.val}
		c.y += m.val
	case down:
		mv.y = cRange{c.y, c.y - m.val}
		c.y -= m.val
	default:
		panic(fmt.Errorf("invalid movement: %v", m.dir))
	}

	return mv
}

func main() {

	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	inputs, err := util.ReadLines("input.txt")
	util.Must(err)

	movs1 := computeCableMovements(inputs[0])
	movs2 := computeCableMovements(inputs[1])

	ints := computeIntersections(movs1, movs2)

	var wg sync.WaitGroup
	wg.Add(2)
	go partOne(ints, &wg)
	go partTwo(ints, movs1, movs2, &wg)
	wg.Wait()

	fmt.Println("DONE!")
}

func computeIntersections(movs1, movs2 []movement) []coord {

	coords := []coord{}

	f := func(mov1, mov2 movement) []coord {

		mov1.ascending()
		mov2.ascending()

		res := []coord{}
		for x := mov1.x.start; x <= mov1.x.end; x++ {
			for y := mov1.y.start; y <= mov1.y.end; y++ {

				c := coord{x, y}
				if c.isIntersected(mov2) {
					res = append(res, c)
				}

			}
		}
		return res
	}

	for _, mov1 := range movs1 {
		for _, mov2 := range movs2 {
			coords = append(coords, f(mov1, mov2)...)
		}
	}

	return coords
}

func partOne(ints []coord, wg *sync.WaitGroup) {

	defer wg.Done()

	sort.Slice(ints, func(i, j int) bool {
		return ints[i].manhattanDistance() < ints[j].manhattanDistance()
	})

	fmt.Printf("part one result: %d\n", ints[0].manhattanDistance())

}

// --- Part Two ---

// It turns out that this circuit is very timing-sensitive; you actually need to minimize the signal delay.

// To do this, calculate the number of steps each wire takes to reach each intersection; choose the intersection where the sum of both wires' steps is lowest.
// If a wire visits a position on the grid multiple times, use the steps value from the first time it visits that position when calculating the total value of a specific intersection.

// The number of steps a wire takes is the total number of grid squares the wire has entered to get to that location,
// including the intersection being considered. Again consider the example from above:

// ...........
// .+-----+...
// .|.....|...
// .|..+--X-+.
// .|..|..|.|.
// .|.-X--+.|.
// .|..|....|.
// .|.......|.
// .o-------+.
// ...........

// In the above example, the intersection closest to the central port is reached after 8+5+5+2 = 20 steps by the first wire and
// 7+6+4+3 = 20 steps by the second wire for a total of 20+20 = 40 steps.

// However, the top-right intersection is better: the first wire takes only 8+5+2 = 15 and the second wire takes only 7+6+2 = 15, a total of 15+15 = 30 steps.

// Here are the best steps for the extra examples from above:

//     R75,D30,R83,U83,L12,D49,R71,U7,L72
//     U62,R66,U55,R34,D71,R55,D58,R83 = 610 steps
//     R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
//     U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = 410 steps

// What is the fewest combined steps the wires must take to reach an intersection?

func countSteps(c coord, movs []movement) int {

	var steps int
	for _, mov := range movs {

		if c.isIntersected(mov) {
			stepX, stepY := c.x-mov.x.start, c.y-mov.y.start
			if stepX < 0 {
				stepX = -stepX
			}
			if stepY < 0 {
				stepY = -stepY
			}
			return steps + stepX + stepY
		}

		mov.ascending()
		steps += (mov.x.end - mov.x.start) + (mov.y.end - mov.y.start)
	}
	return steps
}

func countMultipleSteps(c coord, mov1, mov2 []movement) int {
	return countSteps(c, mov1) + countSteps(c, mov2)
}

func partTwo(ints []coord, mov1, mov2 []movement, wg *sync.WaitGroup) {

	defer wg.Done()

	sort.Slice(ints, func(i, j int) bool {

		return countMultipleSteps(ints[i], mov1, mov2) <
			countMultipleSteps(ints[j], mov1, mov2)
	})

	fmt.Printf("part two result: %d\n", countMultipleSteps(ints[0], mov1, mov2))

}
