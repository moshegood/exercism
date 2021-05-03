package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("hamming distance can only be calculated when the strings are the same length")
	}
	diffs := 0
	bs := []rune(b)
	for i, c := range a {
		if c != bs[i] {
			diffs++
		}
	}
	return diffs, nil
}
