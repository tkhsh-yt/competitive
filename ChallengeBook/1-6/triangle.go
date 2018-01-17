package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				d := a[i] + a[j] + a[k]
				if a[i] < a[j]+a[k] && max < d {
					max = d
				}
			}
		}
	}

	fmt.Println(max)
}
