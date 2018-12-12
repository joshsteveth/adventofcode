package util

import (
	"fmt"
	"testing"
)

func TestIsCorrectBoxes(t *testing.T) {
	tt := map[string]struct {
		a             string
		b             string
		maxErr        int
		sharedStr     string
		expectedError error
	}{
		"incorrect length": {"abc", "abcd", 1, "", fmt.Errorf("length of both string must be equal")},
		"too many errors":  {"achde", "abhij", 2, "ah", fmt.Errorf("current error: 3, max error: 2")},
		"correct boxes":    {"abedg", "abcfg", 2, "abg", nil},
	}

	for testname, tc := range tt {
		t.Run(testname, func(t *testing.T) {
			str, err := IsCorrectBoxes(tc.a, tc.b, tc.maxErr)
			if str != tc.sharedStr {
				t.Errorf("shared string: %s is not equal expected: %s", str, tc.sharedStr)
			}

			if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", tc.expectedError) {
				t.Errorf("error %v is not equal expected one: %v", err, tc.expectedError)
			}
		})
	}
}
