package main

import "fmt"

var X, A, B int

func main() {
	fmt.Scanf("%d", &X)
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)

	fmt.Println((X - A) % B)
}
