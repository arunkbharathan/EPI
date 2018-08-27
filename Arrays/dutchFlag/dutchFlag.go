package dutchflag

const f = 20

// func main() {
// 	x := [f]int{-1, 11, 7, 5, 3, 6, 8, 9, 2, 6, 9, 13, 5, 3, 2, 6, 0, 10, 4, 8}
// 	pivot := 9
// 	fmt.Println(x)
// 	fmt.Println(dutch1(x, pivot))
// 	fmt.Println(dutch2(x, pivot))
// }

func dutch1(a [f]int, pivot int) [f]int {
	end := len(a) - 1

	for i := 0; i < end; i++ {
		if a[i] > pivot {
			a[end], a[i] = a[i], a[end]
			i--
			end--

		}
	}

	end = len(a) - 1
	for i := 0; i < end; i++ {
		if a[i] == pivot {
			if a[end] < pivot {
				a[end], a[i] = a[i], a[end]
			} else {
				end--
				i--
			}
		}
	}
	return a
}

func dutch2(a [f]int, pivot int) [f]int {
	s := 0
	l := len(a)
	e := 0
	for e < l {
		if a[e] < pivot {
			a[s], a[e] = a[e], a[s]
			s++
			e++
		} else if a[e] == pivot {
			e++
		} else {
			l--
			a[e], a[l] = a[l], a[e]
		}
	}
	return a
}
