package main

import (
	"fmt"
	"os"
)

func main() {
	snakestring := "Hello World!"
	if len(os.Args) == 3 {
		snakestring = os.Args[2]
	}
	fmt.Println(getSnakeString(snakestring))
}

func getSnakeString(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += string(s[0+(i)*2])
		res += string(s[1+(i)*4])
		res += string(s[2+(i)*4])
		fmt.Println(res)
	}
	return res
}
