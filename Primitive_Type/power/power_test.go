package power

import (
	"math"
	"testing"
)

var testset = []struct {
	x   float64
	y   int
	out float64
}{
	{1.1, 21, 7.4002},
	{3.14, 2, 9.8596},
	{10, 10, 10000000000},
	{2, -4, 0.0625},
	{3, -2, 0.1111},
}

func TestNthpowerof(t *testing.T) {
	for _, aset := range testset {
		out := nthpowerof(aset.x, aset.y)
		out = math.Round(out*10000) / 10000
		if out != aset.out {
			t.Errorf("Expected %v got %v", aset.out, out)
		}
	}
}
