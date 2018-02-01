package main

import "fmt"

var N int
var C, S, F []int

func main() {
	fmt.Scanf("%d", &N)

	C = make([]int, N-1)
	S = make([]int, N-1)
	F = make([]int, N-1)

	for i := 0; i < N-1; i++ {
		fmt.Scanf("%d %d %d", &C[i], &S[i], &F[i])
	}

	ans := make([]int, N)
	for i := 0; i < N-1; i++ {
		for j := i; j < N-1; j++ {
			ans[i] = Max(ans[i], S[j])
			ans[i] += (F[j] - (ans[i]-S[j])%F[j]) % F[j]
			ans[i] += C[j]
		}
	}

	for i := 0; i < N; i++ {
		fmt.Println(ans[i])
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return x
}
