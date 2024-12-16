package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type dir int

const (
	up dir = iota
	right
	down
	left
)

func move(x, y int, d dir) (int, int) {
	switch d {
	case up:
		return x, y - 1
	case right:
		return x + 1, y
	case down:
		return x, y + 1
	case left:
		return x - 1, y
	}
	panic("no way")
}

func do(maze [][]bool, scores [][]map[dir]int, score, x, y int, d dir) {
	sc, ok := scores[y][x][d]
	if ok && sc <= score {
		return
	}
	scores[y][x][d] = score
	nx, ny := move(x, y, d)
	if !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny, d)
	}
	nx, ny = move(x, y, (d+3)%4)
	if !maze[ny][nx] {
		do(maze, scores, score+1001, nx, ny, (d+3)%4)
	}
	nx, ny = move(x, y, (d+1)%4)
	if !maze[ny][nx] {
		do(maze, scores, score+1001, nx, ny, (d+1)%4)
	}
}

func do2(paths [][]bool, scores [][]map[dir]int, score, x, y int, d dir) {
	paths[y][x] = true
	nx, ny := move(x, y, (d+2)%4)
	if scores[ny][nx][d] == score-1 {
		do2(paths, scores, score-1, nx, ny, d)
	}
	if scores[ny][nx][(d+1)%4] == score-1001 {
		do2(paths, scores, score-1001, nx, ny, (d+1)%4)
	}
	if scores[ny][nx][(d+3)%4] == score-1001 {
		do2(paths, scores, score-1001, nx, ny, (d+3)%4)
	}
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var maze [][]bool
	var scores [][]map[dir]int
	var s string
	for scanner.Scan() {
		s = scanner.Text()
		row := make([]bool, len(s))
		for i, ch := range s {
			if ch == '#' {
				row[i] = true
			}
		}
		maze = append(maze, row)
		sc := make([]map[dir]int, len(s))
		for i := range sc {
			sc[i] = make(map[dir]int)
		}
		scores = append(scores, sc)
	}
	do(maze, scores, 0, 1, len(maze)-2, right)
	fmt.Println(scores[1][len(scores[0])-2])

	paths := make([][]bool, len(maze))
	for i := range paths {
		paths[i] = make([]bool, len(maze[i]))
	}
	do2(paths, scores, 93436, len(scores[0])-2, 1, 0)
	sum := 0
	for _, row := range paths {
		for _, x := range row {
			if x {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day16.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	//part2(buf)

}
