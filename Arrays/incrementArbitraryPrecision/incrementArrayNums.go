package main

import (
	"fmt"
)

func main() {
	a := []int{9, 8, 9, 9}
	a = incrementArray(a)
	fmt.Println(a)
}

func incrementArray(a []int) []int {
	c := 1
	l := len(a) - 1
	for c > 0 {
		a[l] += c
		if a[l] == 10 {
			a[l] = 0
			l--
			if l == -1 {
				a = append([]int{1}, a...)
				c = 0
			}
		} else {
			c = 0
		}
	}
	return a
}
