package main

import (
	"fmt"
)

func main() {
	//var a = []int{1, 4, 5, 7, 6, 3, 2, 9}
	//var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = []int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}
	//var a = []int{10, 9, 8, 7}
	//var a = []int{6, 5, 3, 7, 34, 2, 5, 7, 3, 5, 87, 9, 6, 3, 2, 4}
	//var a = []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	i, j, p := findMaxSubArray(a)
	if p > 0 {
		fmt.Printf("At a[%v]=%v and a[%v]=%v profit is %v\n", i, a[i], j, a[j], p)
	} else {
		fmt.Println("Stock is going down.")
	}
	fmt.Printf("Profit %v\n", maxProfit(a))
}

func findMaxSubArray(a []int) (int, int, int) {
	sofari, sofarj, best := 0, 0, 0
	curri, currj, currbest := -1, -1, 0
	j := 0
	for i := 1; i < len(a); i++ {
		j = i - 1

		if currbest+(a[i]-a[j]) > 0 {

			currbest = a[i] - a[j] + currbest
			if j == curri {
				curri = i
			} else {
				curri = i
				currj = j
			}

		} else {

			currbest = 0
			curri = -1
			currj = -1
		}
		if best < currbest {
			best = currbest
			sofari = curri
			sofarj = currj
		}
	}

	return sofari, sofarj, best
}

func maxProfit(a []int) int {
	sofarmin := 99999
	maxprofit := 0
	for _, price := range a {
		maxprofttoday := price - sofarmin
		if maxprofit < maxprofttoday {
			maxprofit = maxprofttoday
		}
		if price < sofarmin {
			sofarmin = price
		}
	}
	return maxprofit
}
