package main

import "fmt"

func main() {
	var N, A, B int

	fmt.Scanf("%d %d %d", &N, &A, &B)

	if (B-A-1)%2 == 0 {
		fmt.Println("Borys")
	} else {
		fmt.Println("Alice")
	}
}
