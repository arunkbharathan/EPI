package div

func divide(x, y uint64) uint64 {

	var quo uint64
	for x >= y {
		var power uint64
		tempY := y
		tempY <<= 1
		for tempY <= x {
			power++
			tempY <<= 1
		}
		x = x - y<<power
		quo += 1 << power
	}

	return quo
}
