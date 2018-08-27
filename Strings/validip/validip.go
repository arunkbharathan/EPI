package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arg string
	flag.StringVar(&arg, "ip", "19216811", "command 19216811")
	flag.Parse()
	a := validips(arg)
	for _, j := range a {

		fmt.Println(j)
	}
	fmt.Println(len(a))
}

func validips(in string) []string {
	var output []string
	for d1 := 1; d1 <= 3; d1++ {
		ip1 := in[:d1] + "." + in[d1:]
		for d2 := 1; d2 <= 3; d2++ {
			ip2 := ip1[:d1+1+d2] + "." + ip1[d1+1+d2:]
			for d3 := 1; d3 <= 3; d3++ {
				if len(ip2[d1+1+d2+1:]) < d3 {
					//	fmt.Println(ip2)
					continue
				}
				// if ip2[d1+1+d2+1+d3:] == "" {
				// 	//	fmt.Println(ip2)
				// 	continue
				// }
				ip3 := ip2[:d1+1+d2+1+d3] + "." + ip2[d1+1+d2+1+d3:]
				strs := strings.Split(ip3, ".")
				for _, j := range strs {
					if j == "" {
						goto X
					}
					val, err := strconv.Atoi(j)
					if err != nil {
						fmt.Println(err)
						goto X
					}
					if val > 255 {
						goto X
					}
				}
				output = append(output, ip3)
			X:
			}

		}
	}
	return output
}
