package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(isSubstr("smx man sxk", "man"))
}

func isSubstr(t, s string) bool {
	if len(t) < len(s) {
		return false
	}
	power := int(math.Pow(26, float64(len(s))-1))
	tHash := 0
	sHash := 0
	for i, j := range s {
		tHash = tHash*26 + int(t[i])
		sHash = sHash*26 + int(j)
	}
	if tHash == sHash && strings.Compare(t[:3], s) == 0 {
		return true
	}
	for i := len(s); i < len(t); i++ {
		tHash = (tHash-(int(t[i-len(s)])*power))*26 + int(t[i])
		//fmt.Println(string(t[i-len(s)]), " ", t[i-len(s)+1:i+1], " ", string(t[i]))
		if tHash == sHash && strings.Compare(t[i-len(s)+1:i+1], s) == 0 {
			return true
		}
	}
	return false
}
