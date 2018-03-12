package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Printf("%s", string([]rune(s)[i]))
	}
	fmt.Println()
}
