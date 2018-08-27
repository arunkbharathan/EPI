package main

import (
	"bytes"
	"fmt"
)

type ll struct {
	data int
	next *ll
}

func (l *ll) String() string {
	h := l.next
	s := ""
	w := bytes.NewBufferString(s)
	for h.next != nil {
		fmt.Fprintf(w, "%d ", h.data)
		h = h.next
	}
	fmt.Fprintf(w, "%d ", h.data)
	fmt.Fprintln(w, "")
	return w.String()
}

func main() {
	a := []int{2, 5, 7}
	b := []int{3, 11}
	x := ll{}
	y := ll{}
	var t = &x
	for _, j := range a {
		t.next = &ll{data: j}
		t = t.next
	}
	t = &y
	for _, j := range b {
		t.next = &ll{data: j}
		t = t.next
	}

	fmt.Println(&x)
	fmt.Println(&y)
	z := merge(&x, &y)
	fmt.Println(z)
}

func merge(x, y *ll) *ll {
	z := ll{}
	var hx = x.next
	var hy = y.next
	var hz = &z
	for hx != nil {

		if hx.data < hy.data {
			hz.next = &ll{data: hx.data}
			hz = hz.next
			hx = hx.next
		} else {
			hz.next = &ll{data: hy.data}
			hz = hz.next
			hy = hy.next
			if hy == nil {
				break
			}
		}

	}
	for hx != nil {
		hz.next = &ll{data: hx.data}
		hz = hz.next
		hx = hx.next
	}
	for hy != nil {
		hz.next = &ll{data: hy.data}
		hz = hz.next
		hy = hy.next
	}
	return &z
}
