package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int

	fmt.Scanf("%d", &N)

	d := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &d[i])
	}

	sort.Sort(sort.IntSlice(d))

	x := 1
	for i := N - 1; i > 0; i-- {
		if d[i] > d[i-1] {
			x++
		}
	}

	fmt.Println(x)
}
