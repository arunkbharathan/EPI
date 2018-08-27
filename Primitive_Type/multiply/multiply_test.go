package mul

import "testing"

var testset = []struct {
	x   uint64
	y   uint64
	out uint64
}{
	{4, 3, 12},
	{10, 10, 100},
	{32, 64, 2048},
	{237, 8273, 1960701},
}

func TestMultiply(t *testing.T) {
	for _, set := range testset {
		out := multiply(set.x, set.y)
		if out != set.out {
			t.Errorf("Expected %v got %v", set.out, out)
		}
	}
}
