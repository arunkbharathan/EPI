package paritycalc

import "testing"

var tests = []struct {
	in  uint64
	out uint16
}{
	{8, 1},
	{3, 0},
	{89347783627, 0},
	{9223372036854775807, 1},
	{922337203685477587, 1},
}

func TestParity1(t *testing.T) {
	for _, testset := range tests {
		out := parity1(testset.in)
		if out != testset.out {
			t.Errorf("Expected parity %v got %v", testset.out, out)
		}
	}
}
func TestParity2(t *testing.T) {
	for _, testset := range tests {
		out := parity2(testset.in)
		if out != testset.out {
			t.Errorf("Expected parity %v got %v", testset.out, out)
		}
	}
}
func TestParity3(t *testing.T) {
	for _, testset := range tests {
		out := parity3(testset.in)
		if out != testset.out {
			t.Errorf("Expected parity %v got %v", testset.out, out)
		}
	}
}

func BenchmarkParity1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testset := range tests {
			out := parity1(testset.in)
			if out != testset.out {
				b.Errorf("Expected parity %v got %v", testset.out, out)
			}
		}
	}
}

func BenchmarkParity2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testset := range tests {
			out := parity2(testset.in)
			if out != testset.out {
				b.Errorf("Expected parity %v got %v", testset.out, out)
			}
		}
	}
}
func BenchmarkParity3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testset := range tests {
			out := parity3(testset.in)
			if out != testset.out {
				b.Errorf("Expected parity %v got %v", testset.out, out)
			}
		}
	}
}
