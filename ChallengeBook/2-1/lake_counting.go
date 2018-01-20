package main

import "fmt"

/*
10 12
W........WW.
.WWW.....WWW
....WW...WW.
.........WW.
.........W..
..W......W..
.W.W.....WW.
W.W.W.....W.
.W.W......W.
..W.......W.
*/

var N, M int
var field [][]bool

func main() {
	fmt.Scanf("%d %d", &N, &M)

	field = make([][]bool, N)

	for i := 0; i < N; i++ {
		field[i] = make([]bool, M)
	}

	for y := 0; y < N; y++ {
		var s string
		fmt.Scanf("%s", &s)
		for x := 0; x < M; x++ {
			field[y][x] = s[x] == 'W'
		}
	}

	ans := 0
	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if field[y][x] {
				dfs(x, y)
				ans++
			}
		}
	}

	fmt.Println(ans)
}

func dfs(x, y int) {
	field[y][x] = false

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy

			if 0 <= nx && nx < M && 0 <= ny && ny < N && field[ny][nx] {
				dfs(nx, ny)
			}
		}
	}
}
