package main

import (
	"flag"
	"fmt"
)

var nummap = []string{
	"0",
	"1",
	"ABC",
	"DEF",
	"GHI",
	"JKL",
	"MNO",
	"PQRS",
	"TUV",
	"WXYZ",
}

var collect []string

func init() {

	collect = make([]string, 0, 1000000)
}

func main() {
	var svar string
	flag.StringVar(&svar, "svar", "911", "phone number")
	flag.Parse()
	allCombi(svar, 0, "")
	for _, val := range collect {
		fmt.Println(val)
	}
	fmt.Println(len(collect))
}

func allCombi(a string, s int, fin string) {
	if len(a) == s {
		collect = append(collect, fin)
		return
	}

	for _, j := range nummap[a[s]-'0'] {
		allCombi(a, s+1, fin+string(j))
	}

}
