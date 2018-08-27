package reversedigit

import (
	"testing"
)

var testset = []struct {
	x int
	y int
}{
	{314, 413},
	{111, 111},
	{-101, -101},
	{-123, -321},
	{7789, 9877},
	{6767, 7676},
}

func TestReverseDigit(t *testing.T) {
	for _, aset := range testset {
		y := reverseDigit(aset.x)
		if aset.y != y {
			t.Errorf("Expected %v got %v", aset.y, y)
		}
	}
}
