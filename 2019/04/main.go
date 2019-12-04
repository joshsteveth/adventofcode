// --- Day 4: Secure Container ---
package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const min, max = 272091, 815432

type isValid func(int) bool

func main() {

	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	parts := []isValid{partOne, partTwo}

	var wg sync.WaitGroup
	wg.Add(len(parts))

	for i, iv := range parts {
		go func(part int, iv isValid) {
			defer wg.Done()
			var sum int
			for i := min; i <= max; i++ {
				if iv(i) {
					sum++
				}
			}
			fmt.Printf("part %d result: %d\n", part, sum)
		}(i+1, iv)
	}

	wg.Wait()

	fmt.Println("DONE!")
}

// You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.

// However, they do remember a few key facts about the password:

//     It is a six-digit number.
//     The value is within the range given in your puzzle input.
//     Two adjacent digits are the same (like 22 in 122345).
//     Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

// Other than the range rule, the following are true:

//     111111 meets these criteria (double 11, never decreases).
//     223450 does not meet these criteria (decreasing pair of digits 50).
//     123789 does not meet these criteria (no double).

// How many different passwords within the range given in your puzzle input meet these criteria?

func partOne(n int) bool {
	var (
		prev     int
		adjacent bool
	)
	for i := 5; i >= 0; i-- {
		denum := int(math.Pow10(i))
		digit := n / denum
		if digit < prev {
			return false
		} else if digit == prev {
			adjacent = true
		}
		prev = digit
		n -= digit * denum
	}

	return adjacent
}

// --- Part Two ---

// An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

// Given this additional criterion, but still ignoring the range rule, the following are now true:

//     112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
//     123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
//     111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).

// How many different passwords within the range given in your puzzle input meet all of the criteria?

func partTwo(n int) bool {

	var (
		prev, streakLength int
		onStreak, adjacent bool
	)

	for i := 5; i >= 0; i-- {
		denum := int(math.Pow10(i))
		digit := n / denum
		if digit < prev {
			return false
		}

		if adjacent {
			goto assignDigit
		}

		if digit == prev {
			if !onStreak {
				streakLength = 2
			} else {
				streakLength++
			}

			onStreak = true
			goto assignDigit
		}
		if onStreak && streakLength == 2 {
			adjacent = true
		}

		onStreak = false
		streakLength = 0

	assignDigit:
		prev = digit
		n -= digit * denum
	}

	// check for the last combination
	if onStreak && streakLength == 2 {
		adjacent = true
	}

	return adjacent
}
