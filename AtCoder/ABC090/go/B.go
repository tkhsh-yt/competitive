package main

import "fmt"

var N int
var M int

func main() {
	fmt.Scanf("%d %d", &N, &M)

	fmt.Println(N, M)

	flip := (N - 1) * (M - 1)

	fmt.Println(flip)
}
