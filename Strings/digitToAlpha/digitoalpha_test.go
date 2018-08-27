package main

import "testing"

func Benchmark_allCombi(b *testing.B) {
	for i := 0; i < b.N; i++ {

		allCombi("9746637338", 0, "")
	}

}
