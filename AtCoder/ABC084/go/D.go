package main

import "fmt"

var Q int
var l, r []int

func main() {
	fmt.Scanf("%d", &Q)

	l = make([]int, Q)
	r = make([]int, Q)

	for i := 0; i < Q; i++ {
		fmt.Scanf("%d %d", &l[i], &r[i])
	}

	primes := Primes(100001)
	cs := make([]int, 100001)

	for i := 1; i < 100001; i++ {
		if primes[i] && primes[(i+1)/2] {
			cs[i]++
		}
		cs[i] += cs[i-1]
	}

	for i := 0; i < Q; i++ {
		fmt.Println(cs[r[i]+1] - cs[l[i]-1])
	}
}

func Primes(n int) []bool {
	primes := make([]bool, n)

	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		for j := 2; i*j < n; j++ {
			primes[i*j] = false
		}
	}

	return primes
}
