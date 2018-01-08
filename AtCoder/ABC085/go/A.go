package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Scanf("%s", &str)
	new := []rune(str)
	new[3] = '8'
	fmt.Printf(string(new))
}
