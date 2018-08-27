package palindrome

import (
	"math"
)

func palindrome(x int) bool {
	for x != 0 {

		e := math.Log10(float64(x))
		l := x / int(math.Pow10(int(e)))
		r := int(x) % 10
		if l != r {
			return false
		}
		x -= int(math.Pow10(int(e)) * float64(l))
		x /= 10

	}
	return true
}
