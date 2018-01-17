package main

import "fmt"

func main() {
	var L, n int

	fmt.Scanf("%d %d", &L, &n)

	x := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x[i])
	}

	min := 0
	max := 0
	for i := 0; i < n; i++ {
		if x[i] > L-x[i] {
			if L-x[i] > min {
				min = L - x[i]
			}
			if x[i] > max {
				max = x[i]
			}
		} else {
			if x[i] > min {
				min = x[i]
			}
			if L-x[i] > max {
				max = L - x[i]
			}
		}
	}
	fmt.Println(min, max)
}
