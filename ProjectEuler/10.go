package main

import "fmt"

func main() {
	n := 2000000
	ps := primes(n)

	ans := 0
	for i := 0; i < n; i++ {
		if ps[i] {
			ans += i
		}
	}

	fmt.Println(ans)
}

func primes(n int) []bool {
	ps := make([]bool, n)

	for i := 2; i < n; i++ {
		ps[i] = true
	}

	for i := 2; i*i <= n; i++ {
		for j := 2; i*j < n; j++ {
			ps[i*j] = false
		}
	}

	return ps
}
