package main

import (
	"testing"
)

func TestDecodeInput(t *testing.T) {
	tt := map[string]struct {
		s       string
		isError bool
		c       claim
	}{
		"not splittable by @":           {"laekmglamkef", true, claim{}},
		"not splittable by :":           {"#123 @ foo", true, claim{}},
		"not splittable by ,":           {"#123 @ foo : bar", true, claim{}},
		"left padding not convertable":  {"#123 @ foo,2 : bar", true, claim{}},
		"right padding not convertable": {"#123 @ 3, foo : bar", true, claim{}},
		"not splittable by x":           {"#123 @ 3,2 : bar", true, claim{}},
		"width not convertable":         {"#123 @ 3,2 : barx4", true, claim{}},
		"height not convertable":        {"#123 @ 3,2 : 5xbar", true, claim{}},
		"correct case": {"#123 @ 3,2: 5x4", false, claim{
			id:          "123",
			leftPadding: 3,
			topPadding:  2,
			width:       5,
			height:      4,
		}},
	}

	for testname, tc := range tt {
		t.Run(testname, func(t *testing.T) {
			c, err := decodeInput(tc.s)
			if tc.isError {
				if err == nil {
					t.Errorf("error is expected in this test case")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if c != tc.c {
				t.Errorf("unexpected result. \ngot %+v \nwanted %+v", c, tc.c)
			}
		})
	}
}

func TestPopulateMap(t *testing.T) {
	m := make(map[int]map[int]int)

	claim, err := decodeInput("#123 @ 3,2: 5x4")
	if err != nil {
		t.Fatalf("failed to decode input: %v", err)
	}

	populateMap(m, claim)

	// so (2,3) until (5,7) should be 1
	// the rest should be 0

	for i := 2; i < 6; i++ {
		for j := 3; j < 8; j++ {
			if m[i][j] != 1 {
				t.Errorf("expected a 1 for position (%d,%d)", i, j)
			}
		}
	}

	//try 2nd case for 3 different claims
	cases := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	m2 := make(map[int]map[int]int)

	for _, c := range cases {
		claim, err := decodeInput(c)
		if err != nil {
			t.Fatalf("failed to decode input: %v", err)
		}

		populateMap(m2, claim)
	}

	// so (3,3) until (4,4) should be 2

	for i := 3; i < 5; i++ {
		for j := 3; j < 5; j++ {
			if m2[i][j] != 2 {
				t.Errorf("expected a 1 for position (%d,%d)", i, j)
			}
		}
	}
}
