package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

type node struct {
	name string
	prev *node
	next []*node
}

const input = "input.txt"

func main() {

	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	inputs, err := util.ReadLines(input)
	util.Must(err)

	availableNodes := buildNodes(inputs)

	var wg sync.WaitGroup
	wg.Add(2)
	go partOne(availableNodes, &wg)
	go partTwo(availableNodes, &wg)
	wg.Wait()

	fmt.Println("DONE!")

}

func buildNodes(inputs []string) map[string]*node {
	availableNodes := make(map[string]*node)

	for _, inp := range inputs {
		back, front := splitInput(inp)

		backNode, backOK := availableNodes[back]
		frontNode, frontOK := availableNodes[front]

		switch {
		case backOK && frontOK:
			frontNode.prev = backNode
			backNode.next = append(backNode.next, frontNode)
		case backOK:
			fn := &node{name: front, prev: backNode}
			availableNodes[front] = fn
			backNode.next = append(backNode.next, fn)
		case frontOK:
			bn := &node{name: back, next: []*node{frontNode}}
			availableNodes[back] = bn
			frontNode.prev = bn
		default:
			fn := &node{name: front}
			bn := &node{name: back}
			bn.next = []*node{fn}
			fn.prev = bn
			availableNodes[back] = bn
			availableNodes[front] = fn
		}
	}
	return availableNodes
}

func splitInput(s string) (a, b string) {
	str := strings.Split(s, ")")
	return str[0], str[1]
}

func partOne(availableNodes map[string]*node, wg *sync.WaitGroup) {
	defer wg.Done()

	var totalOrbits int

	nodeOrbits := make(map[string]int)

	for name, node := range availableNodes {

		var orbits int

		activeNode := node

	nodeLoop:
		for {

			prevNode := activeNode.prev
			if prevNode == nil {
				break nodeLoop
			}

			if v, ok := nodeOrbits[prevNode.name]; ok {
				orbits += v + 1
				break nodeLoop
			}

			orbits++
			activeNode = prevNode
		}

		totalOrbits += orbits
		nodeOrbits[name] = orbits
	}

	fmt.Printf("part one answer: %d\n", totalOrbits)
}

func partTwo(availableNodes map[string]*node, wg *sync.WaitGroup) {
	defer wg.Done()

	you := availableNodes["YOU"]
	san := availableNodes["SAN"]

	// register all routes of san
	var i int
	sanRoutes := make(map[string]int)
	activeNode := san
	for activeNode.prev != nil {
		activeNode = activeNode.prev
		sanRoutes[activeNode.name] = i
		i++
	}

	// countdown from you, find the same node
	i = 0
	activeNode = you
	for {
		activeNode = activeNode.prev
		if v, ok := sanRoutes[activeNode.name]; ok {
			fmt.Printf("part two answer: %d\n", v+i)
			return
		}
		i++
	}

}
