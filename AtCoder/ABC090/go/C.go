package main

import "fmt"

var N int
var M int

func main() {
	fmt.Scanf("%d %d", &N, &M)

	flip := (N - 2) * (M - 2)

	fmt.Println(Abs(flip))
}

func Abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
