package main

import "fmt"

var a, b int

func main() {

	fmt.Scanf("%d %d", &a, &b)

	if (a*b)%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
