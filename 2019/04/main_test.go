package main

import "testing"

func Test_partTwo(t *testing.T) {

	vals := []int{112233, 123444, 111122}

	for _, v := range vals {
		t.Log(partTwo(v))
	}

}
