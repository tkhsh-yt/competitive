package main

import "fmt"

var N, K int
var x, y []int
var c []string

func main() {
	fmt.Scanf("%d %d", &N, &K)

	x := make([]int, N)
	y := make([]int, N)
	c := make([]string, N)

	K2 := K * 2

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d %s", &x[i], &y[i], &c[i])

		if c[i] == "W" {
			y[i] += K
		}

		x[i] %= K2
		y[i] %= K2
	}

	m := MakeMatrix(K2+1, K2+1)

	for i := 0; i < N; i++ {
		m[y[i]+1][x[i]+1]++
	}

	for y := 1; y < K2+1; y++ {
		for x := 1; x < K2+1; x++ {
			m[y][x] += m[y][x-1] + m[y-1][x] - m[y-1][x-1]
		}
	}

	max := 0
	for y := 0; y < K+1; y++ {
		for x := 0; x < K+1; x++ {
			tmp := m[y][x] - m[y][0] - m[0][x] + m[0][0]
			tmp += m[y][K2] - m[y][x+K] - m[0][K2] + m[0][x+K]
			tmp += m[y+K][x+K] - m[y+K][x] - m[y][x+K] + m[y][x]
			tmp += m[K2][x] - m[K2][0] - m[y+K][x] + m[y+K][0]
			tmp += m[K2][K2] - m[K2][x+K] - m[y+K][K2] + m[y+K][x+K]

			tmp = Max(tmp, N-tmp)
			max = Max(max, tmp)
		}
	}

	fmt.Println(max)
}

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
