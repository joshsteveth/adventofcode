package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	numPlayer  = 403
	maxMarble  = 71920
	multiplier = 100
)

func main() {

	marbles := []int{0}
	players := make([]int, numPlayer)

	currentMarble := 0

	var activePlayer int

	now := time.Now()
	for p := 1; p <= maxMarble*multiplier; p++ {
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

		if p%maxMarble == 0 {
			fmt.Printf("[%d] time lapsed: %v\n", p/maxMarble, time.Since(now))
			now = time.Now()
		}
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

// my stupid algorithm result:
/*
[1] time lapsed: 1.752024315s
[2] time lapsed: 6.252138579s
[3] time lapsed: 13.819571377s
[4] time lapsed: 19.144302372s
[5] time lapsed: 18.335451464s
[6] time lapsed: 25.76408986s
[7] time lapsed: 56.870075741s
[8] time lapsed: 1m8.804038773s
[9] time lapsed: 1m6.613367449s
[10] time lapsed: 1m5.907787067s
[11] time lapsed: 1m5.957266656s
[12] time lapsed: 1m4.999713445s
[13] time lapsed: 57.242029387s
[14] time lapsed: 1m49.828245797s
[15] time lapsed: 1m47.525867552s
[16] time lapsed: 1m47.819043789s
[17] time lapsed: 1m46.264084301s
[18] time lapsed: 1m43.857241953s
[19] time lapsed: 1m44.898912324s
[20] time lapsed: 1m41.522474966s
[21] time lapsed: 1m39.602383333s
[22] time lapsed: 1m36.660105396s
[23] time lapsed: 1m41.803168518s
[24] time lapsed: 1m40.672783788s
[25] time lapsed: 1m40.694895464s
[26] time lapsed: 1m44.352048841s
[27] time lapsed: 1m36.927718347s
[28] time lapsed: 1m35.970024372s
[29] time lapsed: 1m35.101491065s
[30] time lapsed: 1m34.743830195s
[31] time lapsed: 2m35.146695099s
[32] time lapsed: 3m43.66332305s
[33] time lapsed: 3m58.049918446s
[34] time lapsed: 4m3.149531663s
[35] time lapsed: 4m7.925544436s
[36] time lapsed: 3m44.702378735s
[37] time lapsed: 3m44.91539642s
[38] time lapsed: 3m45.016473295s
[39] time lapsed: 3m47.861672581s
[40] time lapsed: 3m59.344905391s
[41] time lapsed: 3m41.521830833s
[42] time lapsed: 3m35.930557373s
[43] time lapsed: 3m43.657859717s
[44] time lapsed: 3m36.811961865s
[45] time lapsed: 3m40.04405513s
[46] time lapsed: 3m38.912628089s
[47] time lapsed: 3m36.474812863s
[48] time lapsed: 3m42.739152054s
[49] time lapsed: 3m42.206371564s
[50] time lapsed: 3m36.637187637s
[51] time lapsed: 3m26.590868557s
[52] time lapsed: 3m26.682058472s
[53] time lapsed: 3m26.461565267s
[54] time lapsed: 3m26.675567355s
[55] time lapsed: 3m38.370765944s
[56] time lapsed: 3m34.200868905s
[57] time lapsed: 3m27.481630346s
[58] time lapsed: 3m27.678042531s
[59] time lapsed: 3m27.440518547s
[60] time lapsed: 3m27.020456999s
[61] time lapsed: 3m26.538017215s
[62] time lapsed: 3m26.750922256s
[63] time lapsed: 3m26.842887466s
[64] time lapsed: 3m27.309797152s
[65] time lapsed: 3m26.965490778s
[66] time lapsed: 3m27.477439933s
[67] time lapsed: 3m27.47420453s
[68] time lapsed: 3m27.180874598s
[69] time lapsed: 3m25.557736632s
[70] time lapsed: 3m26.207997803s
[71] time lapsed: 4m55.503872137s
[72] time lapsed: 7m13.071586632s
[73] time lapsed: 7m14.597258176s
[74] time lapsed: 7m17.119501817s
[75] time lapsed: 7m18.149805384s
[76] time lapsed: 7m18.494019506s
[77] time lapsed: 7m17.657308505s
[78] time lapsed: 7m19.383124454s
[79] time lapsed: 7m20.639700942s
[80] time lapsed: 7m20.367351442s
[81] time lapsed: 7m22.44644308s
[82] time lapsed: 7m22.865784206s
[83] time lapsed: 7m26.331478321s
[84] time lapsed: 7m25.165052222s
[85] time lapsed: 7m27.817534147s
[86] time lapsed: 7m25.319234361s
[87] time lapsed: 7m23.994309832s
[88] time lapsed: 7m24.259416884s
[89] time lapsed: 7m29.171326109s
[90] time lapsed: 7m31.749917718s
[91] time lapsed: 7m31.870730721s
[92] time lapsed: 7m33.681586246s
[93] time lapsed: 7m34.91153633s
[94] time lapsed: 7m35.530504258s
[95] time lapsed: 7m34.859651758s
[96] time lapsed: 7m35.36183434s
[97] time lapsed: 7m36.706885642s
[98] time lapsed: 7m37.430335836s
[99] time lapsed: 7m40.05557164s
[100] time lapsed: 7m40.250192596s
Highest score: 3668541094
*/
