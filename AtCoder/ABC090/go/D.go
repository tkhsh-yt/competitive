package main

import (
	"fmt"
)

var N, K int

func main() {
	fmt.Scanf("%d %d", &N, &K)

	if K == 0 {
		fmt.Println(N * N)
		return
	}

	ans := 0
	for b := K + 1; b <= N; b++ {
		ans += (N / b) * (b - K)
		ans += Max(N%b-K+1, 0)
	}

	fmt.Println(ans)
}

func Max(x, y int) int {
	if x > 0 {
		return x
	}
	return y
}
