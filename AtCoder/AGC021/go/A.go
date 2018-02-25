package main

import (
	"fmt"
)

var N string

func main() {
	fmt.Scanf("%s", &N)

	cs := []rune(N)

	for i := 0; i < len(cs); i++ {
		c0 := int(cs[i] - '0')
		if i < len(cs)-1 {
			c1 := int(cs[i+1] - '0')
			if (c0-1) + 9 > c0+c1 {
				cs[i] = rune((c0-1) + '0')
				for j := i+1; j < len(cs); j++ {
					cs[j] = '9'
				}
				break
			}
		}
	}

		ans := 0
	for i := 0; i < len(cs); i++ {
		ans += int(cs[i] - '0')
	}

	fmt.Println(ans)
}

