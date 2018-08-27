package reverse

var revarr [1 << 16]uint16

func init() {
	for val := 0; val < len(revarr); val++ {
		index := val
		for i := uint8(0); i < 8; i++ {
			val = int(swapbits(uint16(val), i, 15-i))
		}
		revarr[index] = uint16(val)
		val = index

	}
}
func bitReverse(val uint64) uint64 {
	y1 := val >> 48 & 0xffff
	y2 := val >> 32 & 0xffff
	y3 := val >> 16 & 0xffff
	y4 := val & 0xffff
	ry1 := revarr[y1]
	ry2 := revarr[y2]
	ry3 := revarr[y3]
	ry4 := revarr[y4]
	retval := uint64(ry4)<<48 | uint64(ry3)<<32 | uint64(ry2)<<16 | uint64(ry1)
	return retval
}

func swapbits(val uint16, i, j uint8) uint16 {
	maski := uint16(1)
	maskj := uint16(1)
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

func exor(a, b uint16) bool {
	x := a > 0
	y := b > 0
	return (!x && y) || (x && !y)
}
