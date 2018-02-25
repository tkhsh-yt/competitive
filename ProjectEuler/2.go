package main

import "fmt"

func main() {
	a, b := int64(1), int64(2)
	ans := int64(0)
	for a <= 4000000 {
		if a%2 == 0 {
			ans += a
		}
		c := a + b
		a = b
		b = c
	}

	fmt.Println(ans)
}
