package main

import "fmt"

var n, k int
var a []int

func main() {
	fmt.Scanf("%d %d", &n, &k)

	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	if dfs(0, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(i, sum int) bool {
	if i == n {
		return sum == k
	}

	if dfs(i+1, sum) {
		return true
	}

	if dfs(i+1, sum+a[i]) {
		return true
	}

	return false
}
