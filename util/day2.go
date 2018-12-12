package util

import (
	"fmt"
)

// IsCorrectBoxes return error if the differences between a and b is greater than maxErr
// e.g. a = "abc" and b = "abd"
// if maxErr = 1 then it will return "ab" and nil
// if maxErr = 0 then it will return "" and error
func IsCorrectBoxes(a, b string, maxErr int) (string, error) {
	if len(a) != len(b) {
		return "", fmt.Errorf("length of both string must be equal")
	}

	numErr := 0
	var newStr string

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			numErr++
			if numErr > maxErr {
				return newStr, fmt.Errorf("current error: %d, max error: %d", numErr, maxErr)
			}
			continue
		}

		newStr += string(a[i])
	}

	return newStr, nil
}
