/*

10 10
#S######.#
......#..#
.#.##.##.#
.#........
##.##.####
....#....#
.#######.#
....#.....
.####.###.
....#...G#

*/

package main

import (
	"container/list"
	"fmt"
	"strings"
)

var N, M int
var maze [][]byte

var sx, sy int
var gx, gy int

var d [][]int

func main() {
	fmt.Scanf("%d %d", &N, &M)

	maze = make([][]byte, N)

	d = make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([]int, M)
	}

	for y := 0; y < N; y++ {
		var s string
		fmt.Scanf("%s", &s)
		maze[y] = []byte(s)

		p := strings.Index(s, "S")
		if p > -1 {
			sx = p
			sy = y
		}

		p = strings.Index(s, "G")
		if p > -1 {
			gx = p
			gy = y
		}
	}

	fmt.Println(dfs())
}

type P struct {
	x int
	y int
}

var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

func dfs() int {
	que := list.New()

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			d[i][j] = -1
		}
	}

	que.PushBack(P{0, 0})
	d[sy][sx] = 0

	for que.Len() != 0 {
		p := que.Remove(que.Front()).(P)
		if p.x == gx && p.y == gy {
			break
		}

		for i := 0; i < 4; i++ {
			nx, ny := p.x+dx[i], p.y+dy[i]

			if 0 <= nx && nx < M && 0 <= ny && ny < N &&
				maze[ny][nx] != '#' && d[ny][nx] == -1 {
				que.PushBack(P{nx, ny})
				d[ny][nx] = d[p.y][p.x] + 1
			}
		}
	}

	return d[gy][gx]
}
