package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if ans < i*j && isPalindromic(strconv.Itoa(i*j)) {
				ans = i * j
			}
		}
	}

	fmt.Println(ans)
}

func isPalindromic(s string) bool {
	cs := []rune(s)

	for i := 0; i < len(cs)/2; i++ {
		if cs[i] != cs[len(cs)-i-1] {
			return false
		}
	}

	return true
}
