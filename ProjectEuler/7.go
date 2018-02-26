package main

import "fmt"

func main() {
	max := 1000000
	ps := primes(max)

	ans := 0
	n := 0
	for i := 0; i < max; i++ {
		if ps[i] {
			n++
			if n == 10001 {
				ans = i
			}
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
