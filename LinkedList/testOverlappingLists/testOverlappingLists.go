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
		//fmt.Printf("%d ", h.data)
		fmt.Fprintf(w, "%d ", h.data)
		h = h.next
	}
	fmt.Fprintf(w, "%d ", h.data)
	fmt.Fprintln(w, "")
	return w.String()
}
func main() {
	a := []int{2, 3, 5, 7, 11, 13, 16, 19, 20, 21}
	b := []int{1, 4, 6, 8}
	x := ll{}
	y := ll{}
	var t = &x
	c := &x
	for i, j := range a {
		t.next = &ll{data: j}
		t = t.next
		if i == 3 {
			c = t
		}
	}
	fmt.Println(&x)
	t = &y
	for _, j := range b {
		t.next = &ll{data: j}
		t = t.next
	}
	t.next = c

	fmt.Println(&x)
	fmt.Println(&y)

	z := x.getFirstCycleNode(&y)
	fmt.Println("Overlappe Node data: ", z.data)
}

func (l *ll) getFirstCycleNode(ly *ll) *ll {
	if tf, xc, yc := getOverlapp(l, ly); tf {
		tx, ty := l, ly
		if sm := min(xc, yc); sm == yc && sm == xc {
			tx, ty = l, ly
			for tx != nil {
				tx = tx.next
				ty = ty.next
				if tx == ty {
					return &ll{data: tx.data}
				}
			}
		} else {
			if max(xc, yc) == 1 {
				tx = l
				cnt := xc - sm
				for i := cnt; i > 0; i-- {
					tx = tx.next
				}
			} else {
				ty = ly
				cnt := yc - sm
				for i := cnt; i > 0; i-- {
					ty = ty.next
				}
			}
			for tx != nil {
				tx = tx.next
				ty = ty.next
				if tx == ty {
					return &ll{data: tx.data}
				}
			}
		}
	}

	return new(ll)
}
func getOverlapp(lx, ly *ll) (bool, int, int) {
	xl := lx
	yl := ly
	xc, yc := 0, 0
	for lx.next != nil {
		lx = lx.next
		xc++
	}
	xl = lx
	for ly.next != nil {
		ly = ly.next
		yc++
	}
	yl = ly
	return xl == yl, xc, yc
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return 1
	}
	return 2
}
