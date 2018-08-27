package swap

import "testing"

var testset = []struct {
	val uint64
	i   uint8
	j   uint8
	out uint64
}{
	{40960, 3, 13, 32776},
	{549755813888, 34, 10, 549755813888},
	{68719476752, 36, 4, 68719476752},
	{268435473, 4, 22, 272629761},
}

func BenchmarkSwapbits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, aset := range testset {
			out := swapbits(aset.val, aset.i, aset.j)
			if out != aset.out {
				b.Errorf("Expected %v got %v", aset.out, out)
			}

		}
	}
}
