package main

import (
	"fmt"
)

var N int
var x, y, t []int

func main() {
	fmt.Scanf("%d", &N)

	x = make([]int, N+1)
	y = make([]int, N+1)
	t = make([]int, N+1)

	for i := 1; i < N+1; i++ {
		fmt.Scanf("%d %d %d", &t[i], &x[i], &y[i])
	}

	if travel(0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func travel(i int) bool {
	if i+1 < N+1 {
		dt, dx, dy := t[i+1]-t[i], Abs(x[i]-x[i+1]), Abs(y[i]-y[i+1])
		if dt >= dx+dy && dt%2 == (dx+dy)%2 {
			return travel(i + 1)
		}
		return false
	}

	return true
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
