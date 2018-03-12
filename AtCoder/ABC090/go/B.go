package main

import (
	"fmt"
	"strconv"
)

var A int
var B int

func main() {
	fmt.Scanf("%d %d", &A, &B)

	ans := 0
	for i := A; i <= B; i++ {
		s := []rune(strconv.Itoa(i))
		for j := 0; j < len(s); j++ {
			if s[j] != s[len(s)-j-1] {
				break
			}
			if j == len(s)-1 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
