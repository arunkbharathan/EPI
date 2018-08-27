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
	a := []int{2, 3, 5, 7, 11, 13, 16, 19}
	x := ll{}
	var t = &x
	for _, j := range a {
		t.next = &ll{data: j}
		t = t.next
	}

	fmt.Println(&x)

	x.reverseSublist(1, 5)
	fmt.Println(&x)
}

func (l *ll) reverseSublist(s, e int) {
	cnt := 0
	t := l
	var sh *ll
	for cnt <= s {
		sh = t
		t = t.next
		cnt++
	}
	ss := t
	for s < e {
		s++
		temp := ss.next
		ss.next = temp.next
		temp.next = sh.next
		sh.next = temp

	}

}
