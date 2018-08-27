package main

import (
	"math/rand"
)

func main() {
	ch := make(chan uint64, 5)
	for i := 0; i < 10000000; i++ {
		go func() {
			randab(100, 10000000, ch)
		}()
		//	fmt.Println(<-ch)
		<-ch
	}
}

func randab(a, b uint64, ch chan<- uint64) {
	rend := b - a + 1
	index := uint64(1)
	var result = uint64(1 << 63)
	for result > rend {
		result = 0
		for i := uint64(0); (index << i) < rend; i++ {
			result <<= 1
			result |= uint64(rand.Intn(2))
		}
	}
	ch <- (result + a)

	return
}
