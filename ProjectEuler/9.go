package main

import "fmt"

func main() {
	ans := 0
	for a := 1; a < 1000; a++ {
		for b := a + 1; a+b < 1000; b++ {
			for c := b + 1; a+b+c <= 1000; c++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					ans = a * b * c
					break
				}
			}
		}
	}

	fmt.Println(ans)
}
