package reversedigit

func reverseDigit(x int) int {
	r := 0
	for x != 0 {
		ld := x % 10
		x /= 10
		r = 10*r + ld
	}
	return r
}
