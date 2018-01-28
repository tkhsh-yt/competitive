package main

import "fmt"

var A, B, C, X int

func main() {
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &X)

	ans := 0
	for i := 0; i <= A; i++ {
		if 500*i > X {
			break
		}
		for j := 0; j <= B; j++ {
			if 500*i+100*j > X {
				break
			}
			for k := 0; k <= C; k++ {
				if 500*i+100*j+50*k == X {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
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
