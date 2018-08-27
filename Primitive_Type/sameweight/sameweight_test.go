package sameweight

import (
	"testing"
)

var testset = []struct {
	in  uint64
	out uint64
}{
	{8, 4},
	{2, 1},
	{7, 11},
	{5, 6},
	{8589934592, 4294967296},
}

func TestNearestSameWeight(t *testing.T) {
	for _, aset := range testset {
		out := nearestSameWeight(aset.in)
		if out != aset.out {
			t.Errorf("Expected %v got %v", aset.out, out)
		}
	}
}
