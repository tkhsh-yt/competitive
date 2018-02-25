package main

import "fmt"

func main() {
	ans := 2

	for i := 2; i <= 20; i++ {
		ans *= i / gcd(ans, i)
	}

	fmt.Println(ans)
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
