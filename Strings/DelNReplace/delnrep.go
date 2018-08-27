package delnrep

//relace 1 with 4,4 and remove 2
func delnrep(arr [10]int) [10]int {
	lenaft := 0
	for i, x := range arr {
		switch x {
		case 0:
			break
		case 1:
			lenaft += 2
		case 2:
			_ = lenaft
			arr[i] = 0
		default:
			lenaft++
		}
	}
	end := lenaft - 1
	for i := end; i >= 0; i-- {
		switch arr[i] {
		case 0:
			continue
		default:
			if end > i && arr[end] == 0 {
				arr[end] = arr[i]
				arr[i] = 0
				end--
			}
		}
	}
	lz := 0
	for i := 0; i < lenaft; i++ {
		switch arr[i] {
		case 0:
			continue
		case 1:
			arr[lz] = 4
			lz++
			arr[lz] = 4
			lz++
		default:

			arr[lz] = arr[i]
			lz++

		}
	}
	return arr
}
