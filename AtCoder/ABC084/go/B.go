package main

import (
	"fmt"
	"unicode"
)

var A, B int
var S string

func main() {
	fmt.Scanf("%d %d", &A, &B)
	fmt.Scanf("%s", &S)

	ok := true
	if len(S) != A+B+1 {
		ok = false
	} else {
		for i := 0; i < len(S); i++ {
			if i == A {
				if S[i] != '-' {
					ok = false
				}
			} else if !unicode.IsDigit([]rune(S)[i]) {
				ok = false
			}
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
