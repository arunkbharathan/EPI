package main

import (
	"fmt"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

func main() {
	str := "{(}{}{)}"
	fmt.Println(str, " ", paranTest(str))
	str = "{()}{}{)}"
	fmt.Println(str, " ", paranTest(str))
	str = "{()}{}{(())}"
	fmt.Println(str, " ", paranTest(str))
	str = "[[[[(((({{{{}}}}))))]]]]"
	fmt.Println(str, " ", paranTest(str))
	str = "{,},(,),[,]"
	fmt.Println(str, " ", paranTest(str))
	str = "([]){()}"
	fmt.Println(str, " ", paranTest(str))
	str = "[()[]{()()}]"
	fmt.Println(str, " ", paranTest(str))
	str = "{)"
	fmt.Println(str, " ", paranTest(str))
	fmt.Println(str, " ", paranTest(str))
	str = "[()[]{()()"
	fmt.Println(str, " ", paranTest(str))
	str = "}[()[]{()()}"
	fmt.Println(str, " ", paranTest(str))
}

func paranTest(chars string) bool {
	stack := lls.New() // empty
	hash := make(map[rune]rune)
	hash['}'], hash[']'], hash[')'] = '{', '[', '('
	for _, c := range chars {
		switch c {
		case '{', '[', '(':
			stack.Push(c)
		case '}', ']', ')':
			val, ok := stack.Pop()
			if !ok || val.(rune) != hash[c] {
				return false
			}
		default:
			continue
		}
	}
	return stack.Empty()
}
