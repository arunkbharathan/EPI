package main

import (
	"fmt"
)

func main() {
	//var a = []int{1, 4, 5, 7, 6, 3, 2, 9}
	//var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//var a = []int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}
	//var a = []int{10, 9, 8, 7}
	//var a = []int{6, 5, 3, 7, 34, 2, 5, 7, 3, 5, 87, 9, 6, 3, 2, 4}
	//var a = []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	var a = []int{2, 30, 15, 10, 8, 25, 80}
	fmt.Println(find2buySellMaxProfit(a))
}

func find2buySellMaxProfit(a []int) (int, int) {
	//b := []int{0, -1, 2, -4, 3, -4, 6, -1, 2}
	sp := 0
	p := 0
	sum := 0
	disconnectP := 0
	for i, v := range a {
		if i == 0 {
			continue
		}
		d := v - a[i-1]
		sum += d
		if sum < 0 {
			sum = 0
			disconnectP = p
		}
		if sum > p {
			if disconnectP > 0 {
				p, sp = sum, disconnectP
			} else {
				p, sp = sum, p
			}

		} else if sum > sp {
			sp = sum
		}

	}

	return p, sp
}
