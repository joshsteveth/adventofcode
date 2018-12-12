package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joshsteveth/adventofcode/util"
)

const (
	filename string = "input.txt"
)

func main() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	// 1st Part
	// Starting with a frequency of zero,
	// what is the resulting frequency after all of the changes in frequency have been applied?
	freq := int(0)

	for _, l := range lines {
		addFreq(&freq, l)
	}

	fmt.Printf("1st part final result: %d\n", freq)

	// 2nd Part
	// What is the first frequency your device reaches twice?
	// Note that your device might need to repeat its list of frequency changes many times before a duplicate frequency is found,
	// and that duplicates might be found while in the middle of processing the list.

	i := 0

	freq = 0
	duplicateMap := map[int]interface{}{0: nil}

	for {
		val, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatalf("%s can't be converted into int: %v", lines[i], err)
		}

		freq += val

		// check if new frequency is already in the map
		// if yes then break it
		if _, ok := duplicateMap[freq]; ok {
			break
		} else {
			duplicateMap[freq] = nil
		}

		// if not then continue by incrementing i
		// unless if i already >= len(lines) then start from 0
		i++

		if i >= len(lines) {
			i = 0
		}
	}

	fmt.Printf("2nd part final result: %d\n", freq)
}

func addFreq(f *int, s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	*f += i

	return nil
}
