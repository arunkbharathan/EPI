package bitscounter

import "testing"

var tests = []struct {
	in  uint64
	out uint64
}{
	{8, 1},
	{3, 2},
	{89347783627, 22},
	{9223372036854775807, 63},
	{922337203685477587, 31},
}

func TestCountBits1(t *testing.T) {
	for _, testset := range tests {
		out := countBits1(testset.in)
		if out != testset.out {
			t.Errorf("Expected no: of bits %v got %v", testset.out, out)
		}
	}
}

func BenchmarkCountBits1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testset := range tests {
			out := countBits1(testset.in)
			if out != testset.out {
				b.Errorf("Expected no: of bits %v got %v", testset.out, out)
			}
		}
	}
}

func TestCountBits2(t *testing.T) {
	for _, testset := range tests {
		out := countBits2(testset.in)
		if out != testset.out {
			t.Errorf("Expected no: of bits %v got %v", testset.out, out)
		}
	}
}

func BenchmarkCountBits2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testset := range tests {
			out := countBits2(testset.in)
			if out != testset.out {
				b.Errorf("Expected no: of bits %v got %v", testset.out, out)
			}
		}
	}
}
