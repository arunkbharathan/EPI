package reverse

import "testing"

var testset = []struct {
	in  uint64
	out uint64
}{
	{0, 0},
	{16777216, 549755813888},
	{2147483648, 4294967296},
	{6442450944, 6442450944},
}

func BenchmarkBitReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, aset := range testset {
			out := bitReverse(aset.in)
			if out != aset.out {
				b.Errorf("Expected %v but got %v", aset.out, out)
			}
		}
	}
}
