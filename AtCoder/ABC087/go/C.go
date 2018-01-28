package main

import "fmt"

var N int
var A [][]int

func main() {
	fmt.Scanf("%d", &N)

	A = MakeMatrix(2, N+1)

	for i := 0; i < 2; i++ {
		for j := 1; j <= N; j++ {
			fmt.Scanf("%d", &A[i][j])
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j <= N; j++ {
			if j != N {
				A[i][j+1] += A[i][j]
			}
		}
	}

	max := 0
	for j := 1; j <= N; j++ {
		tmp := (A[0][j] - A[0][0]) + (A[1][N] - A[1][j-1])
		max = Max(tmp, max)
	}

	fmt.Println(max)
}

// Library

func MakeMatrix(w, h int) [][]int {
	m := make([][]int, w)
	for x := 0; x < w; x++ {
		m[x] = make([]int, h)
	}

	return m
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
