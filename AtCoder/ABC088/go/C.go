package main

import "fmt"

var C [][]int

func main() {

	C = make([][]int, 3)

	for i := 0; i < 3; i++ {
		C[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &C[i][j])
		}
	}

	ans := true
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			ans = ans && (C[i][j]-C[i+1][j]+C[i+1][j+1] == C[i][j+1])
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
