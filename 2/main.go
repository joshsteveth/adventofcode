/*
--- Day 2: Inventory Management System ---

You stop falling through time, catch your breath, and check the screen on the device. "Destination reached. Current Year: 1518. Current Location: North Pole Utility Closet 83N10." You made it! Now, to find those anomalies.

Outside the utility closet, you hear footsteps and a voice. "...I'm not sure either. But now that so many people have chimneys, maybe he could sneak in that way?" Another voice responds, "Actually, we've been working on a new kind of suit that would let him fit through tight spaces like that. But, I heard that a few days ago, they lost the prototype fabric, the design plans, everything! Nobody on the team can even seem to remember important details of the project!"

"Wouldn't they have had enough fabric to fill several boxes in the warehouse? They'd be stored together, so the box IDs should be similar. Too bad it would take forever to search the warehouse for two similar box IDs..." They walk too far away to hear any more.

Late at night, you sneak to the warehouse - who knows what kinds of paradoxes you could cause if you were discovered - and use your fancy wrist device to quickly scan every box and produce a list of the likely candidates (your puzzle input).

To make sure you didn't miss any, you scan the likely candidate boxes again, counting the number that have an ID containing exactly two of any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to get a rudimentary checksum and compare it to what your device predicts.
*/

package main

import (
	"fmt"

	"github.com/joshsteveth/adventofcode/util"
)

const filename string = "input.txt"

/*
For example, if you see the following box IDs:

    abcdef contains no letters that appear exactly two or three times.
    bababc contains two a and three b, so it counts for both.
    abbcde contains two b, but no letter appears exactly three times.
    abcccd contains three c, but no letter appears exactly two times.
    aabcdd contains two a and two d, but it only counts once.
    abcdee contains two e.
    ababab contains three a and three b, but it only counts once.

Of these box IDs, four of them contain a letter which appears exactly twice, and three of them contain a letter which appears exactly three times. Multiplying these together produces a checksum of 4 * 3 = 12.

What is the checksum for your list of box IDs?
*/

func main() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	two, three := 0, 0

	for _, l := range lines {
		runeMap := map[rune]int{}

		for _, r := range l {
			if _, ok := runeMap[r]; !ok {
				runeMap[r] = 1
				continue
			}
			runeMap[r] += 1
		}

		// increment two or three if they are present in the runemap
		for _, val := range runeMap {
			if val == 2 {
				two++
				break
			}
		}

		for _, val := range runeMap {
			if val == 3 {
				three++
				break
			}
		}
	}

	fmt.Printf("1st Part: there are %d two's, %d three's, and the result is: %d\n",
		two, three, two*three)

	maxErr := 1

	// 2nd Part
	// loop through all the lines and find pairs which differ by exactly one character at the same position

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			boxA, boxB := lines[i], lines[j]

			if str, err := util.IsCorrectBoxes(boxA, boxB, maxErr); err == nil {
				fmt.Printf("2nd part: correct boxes is %d and %d with common letters: %s\n",
					i+1, j+1, str)
				return
			}
		}
	}

	fmt.Println("2nd part: no correct boxes found")
}
