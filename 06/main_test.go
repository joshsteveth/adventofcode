package main

import (
	"testing"
)

func Test_manhattanDistance(t *testing.T) {

	tt := []struct {
		c1, c2 *coordinate
		want   float64
	}{
		{&coordinate{1, 1}, &coordinate{4, -3}, 7},
		{&coordinate{1, 1}, &coordinate{3, -1}, 4},
	}

	for _, tc := range tt {

		get := tc.c1.manhattanDistance(tc.c2)

		if get != tc.want {
			t.Errorf("want: %v get: %v", tc.want, get)
		}

	}

}
