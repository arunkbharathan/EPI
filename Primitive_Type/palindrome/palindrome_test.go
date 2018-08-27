package palindrome

import (
	"testing"
)

var testset = []struct {
	x int
	y bool
}{
	{314, false},
	{111, true},
	{17871, true},
	{123321, true},
	{123421, false},
}

func TestPalindrome(t *testing.T) {
	for _, aset := range testset {
		y := palindrome(aset.x)
		if aset.y != y {
			t.Errorf("Expected %v got %v for %v", aset.y, y, aset.x)
		}
	}
}
