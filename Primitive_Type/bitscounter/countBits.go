package bitscounter

func countBits1(x uint64) uint64 {
	var numBits uint64
	for {
		numBits += x & 1
		x = x >> 1
		if x == 0 {
			break
		}
	}
	return numBits
}

func countBits2(x uint64) uint64 {
	var numBits, y uint64
	for {

		y = x &^ (x - 1)
		x = x & (x - 1)
		if y == 0 {
			break
		}
		numBits++
	}
	return numBits
}
