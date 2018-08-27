package div

import "testing"

var testset = []struct {
	x   uint64
	y   uint64
	quo uint64
}{
	{11, 2, 5},
	{20, 2, 10},
	{27, 7, 3},
	{93990923837, 56376, 1667215},
}

func TestDivide(t *testing.T) {
	for _, set := range testset {
		quo := divide(set.x, set.y)
		if quo != set.quo {
			t.Errorf("Expected %v got %v", set.quo, quo)
		}
	}
}
