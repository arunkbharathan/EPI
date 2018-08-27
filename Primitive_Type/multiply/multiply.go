package mul

func multiply(x, y uint64) uint64 {
	var sum uint64
	for x > 0 {
		if (x & 1) > 0 {
			sum = add(sum, y)
		}
		x >>= 1
		y <<= 1
	}
	return sum
}
func add(p, q uint64) uint64 {
	var Cin uint64
	var tot, cnt uint64
	for (p > 0) || (q > 0) {
		A := p & 1
		B := q & 1
		S := fullAdderS(A, B, Cin)
		Cin = fullAdderC(A, B, Cin)
		tot |= S << cnt
		cnt++
		p >>= 1
		q >>= 1
	}
	//if Cin > 0 {
	tot |= Cin << cnt
	//}
	return tot
}

func fullAdderS(A, B, Cin uint64) uint64 {
	return (A & 1) ^ (B & 1) ^ (Cin & 1)
}
func fullAdderC(A, B, Cin uint64) uint64 {
	A = (A & 1)
	B = (B & 1)
	Cin = (Cin & 1)
	return A&B | A&Cin | B&Cin
}
