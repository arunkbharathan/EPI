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
		fmt.Printf("%d ", h.data)
		fmt.Fprintf(w, "%d ", h.data)
		h = h.next
	}
	fmt.Fprintf(w, "%d ", h.data)
	fmt.Fprintln(w, "")
	return w.String()
}
func main() {
	a := []int{2, 3, 5, 7, 11, 13, 16, 19, 20, 21}
	x := ll{}
	var t = &x
	var c = &x
	for i, j := range a {
		t.next = &ll{data: j}
		t = t.next
		if i == 3 {
			c = t
		}
	}
	t.next = c

	//fmt.Println(&x)

	y := x.getFirstCycleNode()
	fmt.Println("Node 1 data: ", y.data)
}

func (l *ll) getFirstCycleNode() *ll {

	var t1, t2 *ll
	t2 = l
	t1 = l
	for i := 0; t2 != nil && t2.next != nil; i += 2 {
		t2 = t2.next.next
		t1 = t1.next
		if t1 == t2 {
			cl := 0
			t1 = t1.next
			for t1.next != t2.next {
				t1 = t1.next
				cl++
			}
			fmt.Println("Cyclicity: ", cl)
			t2 = l
			t1 = l
			for m := 0; m <= cl; m++ {
				t2 = t2.next
			}
			for m := 0; ; m++ {
				t1 = t1.next
				t2 = t2.next
				if t1 == t2 {
					return t1
				}
			}
		}

	}
	return new(ll)
}
