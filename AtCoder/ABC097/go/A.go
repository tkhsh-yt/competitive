package main

import "fmt"

var a, b, c, d int

func main() {
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	if Abs(a-c) <= d || Abs(a-b) <= d && Abs(b-c) <= d {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
