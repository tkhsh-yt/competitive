package main

import (
	"fmt"
	"sort"
)

var N int
var A []int

func main() {
	fmt.Scanf("%d", &N)

	A = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	sort.Ints(A)

	ans := 0
	for i := N - 1; i >= 0; i-- {
		if (N-i)%2 == 1 {
			ans += A[i]
		} else {
			ans -= A[i]
		}
	}

	fmt.Println(ans)
}
