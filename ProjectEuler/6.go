package main

import "fmt"

func main() {
	ds := 0
	for i := 1; i <= 100; i++ {
		ds += i * i
	}

	sd := 0
	for i := 0; i <= 100; i++ {
		sd += i
	}
	sd *= sd

	fmt.Println(sd - ds)
}
