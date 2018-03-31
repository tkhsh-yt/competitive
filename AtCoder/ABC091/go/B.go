package main

import "fmt"

var N, D, X int
var A []int

func main() {
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d %d", &D, &X)

	A = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	ans := X
	for i := 0; i < D; i++ {
		for j := 0; j < N; j++ {
			if i%A[j] == 0 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
