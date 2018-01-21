package main

import (
	"fmt"
	"math"
	"strconv"
)

var a, b string

func main() {
	fmt.Scanf("%s %s", &a, &b)

	ab, _ := strconv.Atoi(a + b)
	c := int(math.Sqrt(float64(ab)))

	if c*c == ab {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
