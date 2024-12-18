package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var lx, ly = 71, 71
var p1len = 1024

func do(maze [][]bool, scores [][]int, score, x, y int) {
	sc := scores[y][x]
	if sc <= score {
		return
	}
	scores[y][x] = score
	nx, ny := x+1, y
	if nx >= 0 && ny >= 0 && nx < lx && ny < ly && !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
	nx, ny = x-1, y
	if nx >= 0 && ny >= 0 && nx < lx && ny < ly && !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
	nx, ny = x, y+1
	if nx >= 0 && ny >= 0 && nx < lx && ny < ly && !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
	nx, ny = x, y-1
	if nx >= 0 && ny >= 0 && nx < lx && ny < ly && !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var s string

	maze := make([][]bool, ly)
	scores := make([][]int, ly)
	for i := range maze {
		maze[i] = make([]bool, lx)
		scores[i] = make([]int, lx)
		for j := range scores[i] {
			scores[i][j] = math.MaxInt
		}
	}
	for i := 0; i < p1len; i++ {
		scanner.Scan()
		s = scanner.Text()
		ss := strings.Split(s, ",")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])
		maze[y][x] = true
	}
	for scanner.Scan() {
	}

	do(maze, scores, 0, 0, 0)
	fmt.Println(scores[ly-1][lx-1])
}

func part2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var s string

	maze := make([][]bool, ly)
	for i := range maze {
		maze[i] = make([]bool, lx)
	}
	i := 0
	for scanner.Scan() {
		fmt.Println(i)
		i++
		scores := make([][]int, ly)
		for i := range maze {
			scores[i] = make([]int, lx)
			for j := range scores[i] {
				scores[i][j] = math.MaxInt
			}
		}

		s = scanner.Text()
		ss := strings.Split(s, ",")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])
		maze[y][x] = true
		do(maze, scores, 0, 0, 0)
		if scores[ly-1][lx-1] == math.MaxInt {
			fmt.Println(x, y)
			break
		}
	}

}

func main() {
	f, err := os.Open("inputs/day18.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
