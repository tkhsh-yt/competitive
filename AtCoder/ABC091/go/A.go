package main

import "fmt"

var A, B, C, D int

func main() {
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &D)

	ans := 0
	if A > B {
		ans += B
	} else {
		ans += A
	}

	if C > D {
		ans += D
	} else {
		ans += C
	}

	fmt.Println(ans)
}
