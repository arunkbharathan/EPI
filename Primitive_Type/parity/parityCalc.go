package paritycalc

func parity1(x uint64) uint16 {
	var p uint16
	for x > 0 {
		p ^= 1
		x = x & (x - 1)
	}
	return p
}

var pt [1 << 16]uint16

func init() {
	for i := 0; i < len(pt); i++ {
		pt[i] = parity1(uint64(i))
	}
}
func parity2(x uint64) uint16 {
	y1 := parity1(x >> 48 & 0xffff)
	y2 := parity1(x >> 32 & 0xffff)
	y3 := parity1(x >> 16 & 0xffff)
	y4 := parity1(x & 0xffff)
	out := pt[y1] ^ pt[y2] ^ pt[y3] ^ pt[y4]
	return out
}

func parity3(x uint64) uint16 {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return uint16(x & 1)
}
