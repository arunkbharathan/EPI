package sameweight

func nearestSameWeight(val uint64) uint64 {
	ele := uint64(0)
	lastone := val & ^(val - 1)
	lastzero := ^val & ^(^val - 1)
	if lastone > lastzero {
		ele = val ^ lastone
		ele = ele ^ (lastone >> 1)
	} else {
		ele = val ^ lastzero
		ele = ele ^ (lastzero >> 1)
	}
	return ele
}
