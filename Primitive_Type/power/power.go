package power

func nthpowerof(x float64, y int) float64 {
	if y < 0 {
		y = -y
		x = 1 / x
	}
	if y == 1 {
		return x
	}
	if y == 0 {
		return 1
	}
	if (y & 1) > 0 {
		return x * nthpowerof(x, y>>1) * nthpowerof(x, y>>1)
	}
	return nthpowerof(x, y>>1) * nthpowerof(x, y>>1)

}
