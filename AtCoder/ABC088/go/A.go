package main

import "fmt"

var N, A int

func main() {
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &A)

	if N%500-A <= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
