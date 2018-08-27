package swap

//does bit swapping of val at index i and j
func Swapbits(val uint64, i, j uint8) uint64 {
	maski := uint64(1)
	maskj := uint64(1)
	maski = 1 << i
	maskj = 1 << j
	res1 := val & maski
	res2 := val & maskj
	if exor(res1, res2) {
		val = val ^ maski
		val = val ^ maskj
	}
	return val
}

func exor(a, b uint64) bool {
	x := a > 0
	y := b > 0
	return (!x && y) || (x && !y)
}
