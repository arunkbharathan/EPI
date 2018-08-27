package dutchflag

import "testing"

var x = [f]int{-1, 11, 7, 5, 3, 6, 8, 9, 2, 6, 9, 13, 5, 3, 2, 6, 0, 10, 4, 8}

func BenchmarkDutch1(b *testing.B) {
	pivot := 6
	for i := 0; i < b.N; i++ {
		y := dutch1(x, pivot)
		if !(y[9] == pivot && y[10] == pivot && y[11] == pivot) {
			b.Errorf("Not Same")
		}
	}
}

func BenchmarkDutch2(b *testing.B) {
	pivot := 6
	for i := 0; i < b.N; i++ {
		y := dutch2(x, pivot)
		if !(y[9] == pivot && y[10] == pivot && y[11] == pivot) {
			b.Errorf("Not Same")
		}
	}
}
