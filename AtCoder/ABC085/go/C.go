package main

import "fmt"

func main() {
	var N, Y int

	fmt.Scanf("%d %d", &N, &Y)

	if Y > 10000*N {
		fmt.Println("-1 -1 -1")
		return
	}

	for x := 0; x <= N; x++ {
		for y := 0; x+y <= N; y++ {
			z := N - (x + y)
			if 10*x+5*y+z == Y/1000 {
				fmt.Printf("%d %d %d", x, y, z)
				return
			}
		}
	}

	fmt.Println("-1 -1 -1")
}
