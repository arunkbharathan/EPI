package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}
type rectangle struct {
	point
	width  float64
	height float64
}

func (R rectangle) String() string {
	return fmt.Sprintf("(x=%v,y=%v) Width:%v Height:%v\n", R.x, R.y, R.width, R.height)
}

func main() {
	R1 := rectangle{point{1, 2}, 5, 6}
	R2 := rectangle{point{2, 3}, 5, 6}
	if doIntersect(R1, R2) {
		fmt.Println("Intersecting Rectangle.")
		fmt.Println(getOverlapingRectangle(R1, R2))
		return
	}
	fmt.Println("They don't intersect.")

}

func doIntersect(R1 rectangle, R2 rectangle) bool {
	return (R1.x <= R2.x && R2.x <= (R1.x+R1.width) &&
		R1.y <= R2.y && R2.y <= (R1.y+R1.height)) || (R2.x <= R1.x && R1.x <= (R2.x+R2.width) &&
		R2.y <= R1.y && R1.y <= (R2.y+R2.height))
}

func getOverlapingRectangle(R1 rectangle, R2 rectangle) *rectangle {
	x := math.Max(R1.x, R2.x)
	y := math.Max(R1.y, R2.y)
	x1 := math.Min(R1.x+R1.width, R2.x+R2.width)
	width := x1 - x
	y1 := math.Min(R1.y+R1.height, R2.y+R2.height)
	height := y1 - y

	return &rectangle{point{x, y}, width, height}
}
