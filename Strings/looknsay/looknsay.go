package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := "1"
	o := looknsay(i, 7)
	fmt.Println(i + " -> " + o)
}

func looknsay(lks string, nth int) string {
	icnt := 1
	var pitm rune
	res := ""
	for i, j := range lks {
		if i == 0 {
			pitm = j
			continue
		}
		if pitm == j {
			icnt++
		} else {
			alphnum := string(pitm)
			cardinal := strconv.Itoa(icnt)
			res = res + cardinal + alphnum
			icnt = 1
		}
		pitm = j
	}
	alphnum := string(pitm)
	cardinal := strconv.Itoa(icnt)
	res = res + cardinal + alphnum
	if nth == 1 {
		return res
	}
	return looknsay(res, nth-1)
}
