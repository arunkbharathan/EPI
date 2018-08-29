package main

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
)

func main() {
	a := []string{"/usr/bin/local", "/usr/lib/../bin/gcc", "scripts//./../scripts/awkscripts/././"}

	for _, x := range a {
		fmt.Println(x, ": ", validPath(x))
	}
	//x := "/usr/lib/../bin/gcc"
	//fmt.Println(x, ": ", validPath(x))
}

func validPath(path string) string {
	stack := arraystack.New()
	//	beg := path[0]
	//tmp :=
	for i, x := range strings.Split(path, "/") {
		switch x {
		case ".":
			continue
		case "..":
			e, k := stack.Pop()
			_, _ = e, k
		case "":
			if i == 0 {
				stack.Push("/")
			}
			continue
		default:

			stack.Push(x)
		}
	}
	out := ""
	stack2 := arraystack.New()

	for {
		if val, ok := stack.Pop(); ok {
			stack2.Push(val.(string))
		} else {
			break
		}
	}
	for {
		if val, ok := stack2.Pop(); ok {
			if val2 := val.(string); val2 != "/" {
				out += val2 + "/"
			} else {
				out += val2
			}

		} else {
			break
		}
	}
	return out
}
