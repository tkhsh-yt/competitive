package main

import "fmt"

func main() {
	n := 600851475143

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			n /= i
		}
	}

	fmt.Println(n)
}
