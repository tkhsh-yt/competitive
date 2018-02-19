package main

import "fmt"

var H, W int
var S []string

func main() {
	fmt.Scanf("%d %d", &H, &W)

	S := make([]string, H)

	for i := 0; i < H; i++ {
		fmt.Scanf("%s", &S[i])
	}

	que := newQueue()
	que.enqueue(&Point{0, 0, 0})

	gone := make([][]bool, H)
	for i := 0; i < H; i++ {
		gone[i] = make([]bool, W)
	}
	gone[0][0] = true

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, -1, 0, 1}

	s := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if ([]rune(S[i]))[j] == '#' {
				s++
			}
		}
	}

	var p *Point
	for !que.isEmpty() {
		p = que.dequeue()

		if p.x == W-1 && p.y == H-1 {
			break
		}

		for i := 0; i < 4; i++ {
			y := p.y + dy[i]
			x := p.x + dx[i]
			if 0 <= x && x < W && 0 <= y && y < H {
				if !gone[y][x] && ([]rune(S[y]))[x] == '.' {
					gone[y][x] = true
					que.enqueue(&Point{x, y, p.m + 1})
				}
			}
		}
	}

	if p.x != W-1 || p.y != H-1 {
		fmt.Println(-1)
	} else {
		fmt.Println(W*H - s - p.m - 1)
	}
}

// Queue

type Point struct {
	x int
	y int
	m int
}

type queue struct {
	container []*Point
}

func newQueue() *queue {
	return &queue{make([]*Point, 0)}
}

func (q *queue) enqueue(s *Point) *Point {
	q.container = append(q.container, s)
	return s
}

func (q *queue) dequeue() *Point {
	s := q.container[0]
	q.container = q.container[1:]
	return s
}

func (q *queue) isEmpty() bool {
	return len(q.container) == 0
}
