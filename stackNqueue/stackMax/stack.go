package main

import (
	"bytes"
	"fmt"
)

type rep struct {
	max   int
	count int
}
type stack struct {
	stack []int
	freq  []rep
}

func (s *stack) top() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	return s.stack[len(s.stack)-1], true
}
func (s *stack) max() (int, bool) {
	if len(s.freq) == 0 {
		return 0, false
	}
	return s.freq[len(s.freq)-1].max, true
}

func (s *stack) push(i int) {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, i)
		s.freq = []rep{{i, 1}}
	} else {
		s.stack = append(s.stack, i)
		if i >= s.freq[len(s.freq)-1].max {
			if i == s.freq[len(s.freq)-1].max {
				s.freq[len(s.freq)-1].count++
			} else {
				s.freq = append(s.freq, rep{i, 1})
			}
		}
	}
}

func (s *stack) pop() (int, bool) {
	if len(s.stack) > 0 {
		i := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		if i == s.freq[len(s.freq)-1].max {
			s.freq[len(s.freq)-1].count--
			if s.freq[len(s.freq)-1].count < 1 {
				s.freq = s.freq[:len(s.freq)-1]
			}
		}
		return i, true
	}
	return 0, false
}

func (s *stack) String() string {
	str := ""
	w := bytes.NewBufferString(str)
	if len(s.stack) > 0 {
		c := 0
		for c = len(s.stack) - 1; c >= 0; c-- {
			if c == len(s.freq)-1 {
				break
			}
			fmt.Fprintf(w, "|%d|\n", s.stack[c])
		}
		for i := len(s.freq) - 1; i >= 0; i-- {
			fmt.Fprintf(w, "|%d| |%d %d|\n", s.stack[c], s.freq[i].max, s.freq[i].count)
			c--
		}
	}
	return w.String()
}

func main() {
	s := &stack{make([]int, 0, 5), make([]rep, 0, 5)}
	fmt.Println(s)
	s.push(2)
	fmt.Println(s)
	s.push(2)
	fmt.Println(s)
	s.push(1)
	fmt.Println(s)
	s.push(4)
	fmt.Println(s)
	s.push(5)
	fmt.Println(s)
	s.push(5)
	fmt.Println(s)
	s.push(3)
	fmt.Println(s)

	s.pop()
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
	s.push(0)
	fmt.Println(s)
	s.push(3)
	fmt.Println(s)

}
