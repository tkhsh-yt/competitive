package main

import "fmt"

func main() {
	var n, m int

	fmt.Scanf("%d %d", &n, &m)

	k := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &k[i])
	}

	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for d := 0; d < n; d++ {
					if k[a]+k[b]+k[c]+k[d] == m {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}

	fmt.Println("No")
}
